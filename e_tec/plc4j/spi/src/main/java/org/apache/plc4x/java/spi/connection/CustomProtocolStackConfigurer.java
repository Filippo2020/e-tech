/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package org.apache.plc4x.java.spi.connection;

import static org.apache.plc4x.java.spi.configuration.ConfigurationFactory.*;

import io.netty.buffer.ByteBuf;
import io.netty.channel.ChannelHandler;
import io.netty.channel.ChannelPipeline;
import io.netty.handler.codec.MessageToMessageCodec;
import org.apache.plc4x.java.api.authentication.PlcAuthentication;
import org.apache.plc4x.java.api.listener.EventListener;
import org.apache.plc4x.java.spi.Plc4xNettyWrapper;
import org.apache.plc4x.java.spi.Plc4xProtocolBase;
import org.apache.plc4x.java.spi.configuration.Configuration;
import org.apache.plc4x.java.spi.context.DriverContext;
import org.apache.plc4x.java.spi.generation.ByteOrder;
import org.apache.plc4x.java.spi.generation.Message;
import org.apache.plc4x.java.spi.generation.MessageInput;

import java.util.List;
import java.util.function.Consumer;
import java.util.function.Function;
import java.util.function.ToIntFunction;

/**
 * Builds a Protocol Stack.
 */
public class CustomProtocolStackConfigurer<C extends Message> implements ProtocolStackConfigurer<C> {

    private final Class<C> basePacketClass;
    private final ByteOrder byteOrder;
    private final Function<Configuration, ? extends Plc4xProtocolBase<C>> protocol;
    private final Function<Configuration, ? extends DriverContext> driverContext;
    private final Function<Configuration, ? extends MessageInput<C>> protocolIO;
    private final Function<Configuration, ? extends ToIntFunction<ByteBuf>> packetSizeEstimator;
    private final Function<Configuration, ? extends Consumer<ByteBuf>> corruptPacketRemover;
    private final MessageToMessageCodec<ByteBuf, ByteBuf> encryptionHandler;

    private final Object[] parserArgs;

    public static <P extends Message> CustomProtocolStackBuilder<P> builder(Class<P> basePacketClass, Function<Configuration, ? extends MessageInput<P>> messageInput) {
        return new CustomProtocolStackBuilder<>(basePacketClass, messageInput);
    }

    /** Only accessible via Builder */
    CustomProtocolStackConfigurer(Class<C> basePacketClass,
                                  ByteOrder byteOrder,
                                  Object[] parserArgs,
                                  Function<Configuration, ? extends Plc4xProtocolBase<C>> protocol,
                                  Function<Configuration, ? extends DriverContext> driverContext,
                                  Function<Configuration, ? extends MessageInput<C>> protocolIO,
                                  Function<Configuration, ? extends ToIntFunction<ByteBuf>> packetSizeEstimator,
                                  Function<Configuration, ? extends Consumer<ByteBuf>> corruptPacketRemover,
                                  MessageToMessageCodec<ByteBuf, ByteBuf> encryptionHandler) {
        this.basePacketClass = basePacketClass;
        this.byteOrder = byteOrder;
        this.parserArgs = parserArgs;
        this.protocol = protocol;
        this.driverContext = driverContext;
        this.protocolIO = protocolIO;
        this.packetSizeEstimator = packetSizeEstimator;
        this.corruptPacketRemover = corruptPacketRemover;
        this.encryptionHandler = encryptionHandler;
    }

    private ChannelHandler getMessageCodec(Configuration configuration) {
        return new GeneratedProtocolMessageCodec<>(basePacketClass, protocolIO.apply(configuration), byteOrder, parserArgs,
            packetSizeEstimator == null ? null : packetSizeEstimator.apply(configuration),
            corruptPacketRemover == null ? null : corruptPacketRemover.apply(configuration));
    }

    /** Applies the given Stack to the Pipeline */
    @Override
    public Plc4xProtocolBase<C> configurePipeline(Configuration configuration, ChannelPipeline pipeline,
                                                                  PlcAuthentication authentication, boolean passive,
                                                                  List<EventListener> ignore) {
        if (this.encryptionHandler != null) {
            pipeline.addLast(this.encryptionHandler);
        }
        pipeline.addLast(getMessageCodec(configuration));
        Plc4xProtocolBase<C> protocol1 = configure(configuration, this.protocol.apply(configuration));
        DriverContext driverContext1 = this.driverContext.apply(configuration);
        if (driverContext1 != null) {
            protocol1.setDriverContext(driverContext1);
        }
        Plc4xNettyWrapper<C> context = new Plc4xNettyWrapper<>(pipeline, passive, protocol1, authentication, basePacketClass);
        pipeline.addLast(context);
        return protocol1;
    }

    /**
     * Used to Build Instances of {@link SingleProtocolStackConfigurer}.
     *
     * @param <BASE_PACKET_CLASS> Type of Created Message that is Exchanged.
     */
    public static final class CustomProtocolStackBuilder<B extends Message> {

        private final Class<B> basePacketClass;
        private final Function<Configuration, ? extends MessageInput<B>> messageInput;
        private Function<Configuration, ? extends DriverContext> driverContext;
        private ByteOrder byteOrder = ByteOrder.BIG_ENDIAN;
        private Object[] parserArgs;
        private Function<Configuration, ? extends Plc4xProtocolBase<B>> protocol;
        private Function<Configuration, ? extends ToIntFunction<ByteBuf>> packetSizeEstimator;
        private Function<Configuration, ? extends Consumer<ByteBuf>> corruptPacketRemover;
        private MessageToMessageCodec<ByteBuf, ByteBuf> encryptionHandler;

        public CustomProtocolStackBuilder(Class<B> basePacketClass, Function<Configuration, ? extends MessageInput<B>> messageInput) {
            this.basePacketClass = basePacketClass;
            this.messageInput = messageInput;
        }

        public CustomProtocolStackBuilder<B> withDriverContext(Function<Configuration, ? extends DriverContext> driverContextClass) {
            this.driverContext = driverContextClass;
            return this;
        }

        public CustomProtocolStackBuilder<B> byteOrder(ByteOrder byteOrder) {
            this.byteOrder = byteOrder;
            return this;
        }

        public CustomProtocolStackBuilder<B> bigEndian() {
            this.byteOrder = ByteOrder.BIG_ENDIAN;
            return this;
        }

        public CustomProtocolStackBuilder<B> littleEndian() {
            this.byteOrder = ByteOrder.LITTLE_ENDIAN;
            return this;
        }

        public CustomProtocolStackBuilder<B> withParserArgs(Object... parserArgs) {
            this.parserArgs = parserArgs;
            return this;
        }

        public CustomProtocolStackBuilder<B> withProtocol(Function<Configuration, ? extends Plc4xProtocolBase<B>> protocol) {
            this.protocol = protocol;
            return this;
        }

        public CustomProtocolStackBuilder<B> withPacketSizeEstimator(Function<Configuration, ? extends ToIntFunction<ByteBuf>> packetSizeEstimator) {
            this.packetSizeEstimator = packetSizeEstimator;
            return this;
        }

        public CustomProtocolStackBuilder<B> withCorruptPacketRemover(Function<Configuration, ? extends Consumer<ByteBuf>> corruptPacketRemover) {
            this.corruptPacketRemover = corruptPacketRemover;
            return this;
        }

        public CustomProtocolStackBuilder<B> withEncryptionHandler(MessageToMessageCodec<ByteBuf, ByteBuf> encryptionHandler) {
            this.encryptionHandler = encryptionHandler;
            return this;
        }

        public CustomProtocolStackConfigurer<B> build() {
            assert this.protocol != null;
            return new CustomProtocolStackConfigurer<>(
                basePacketClass, byteOrder, parserArgs, protocol, driverContext, messageInput, packetSizeEstimator,
                corruptPacketRemover, encryptionHandler);
        }

    }

}
