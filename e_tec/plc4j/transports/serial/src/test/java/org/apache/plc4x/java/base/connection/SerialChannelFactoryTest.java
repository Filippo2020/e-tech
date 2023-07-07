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
package org.apache.plc4x.java.base.connection;

import com.fazecast.jSerialComm.SerialPort;
import io.netty.buffer.ByteBuf;
import io.netty.channel.Channel;
import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.ChannelInitializer;
import io.netty.handler.codec.ByteToMessageCodec;
import org.apache.plc4x.java.api.exceptions.PlcConnectionException;
import org.apache.plc4x.java.transport.serial.DummyHandler;
import org.apache.plc4x.java.transport.serial.SerialChannel;
import org.apache.plc4x.java.transport.serial.SerialChannelFactory;
import org.apache.plc4x.java.transport.serial.SerialChannelHandler;
import org.apache.plc4x.java.transport.serial.SerialSocketAddress;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import static org.junit.jupiter.api.Assertions.assertNull;

import java.net.UnknownHostException;
import java.util.List;

 class SerialChannelFactoryTest {

    private static final Logger logger = LoggerFactory.getLogger(SerialChannelFactoryTest.class);

    @Test
     void showAllPorts() {
        System.out.println("-------------------------------------");
        System.out.println(" Starting to Display all Serial Ports");
        System.out.println("-------------------------------------");
        for (SerialPort commPort : SerialPort.getCommPorts()) {
            System.out.println(commPort.getDescriptivePortName());
            assertNull(commPort);
        }
    }

    @Test
    void createChannel() throws PlcConnectionException, InterruptedException, UnknownHostException {
        SerialChannelFactory asdf = new SerialChannelFactory(new SerialSocketAddress("TEST-port1", DummyHandler.INSTANCE));
        final Channel channel = asdf.createChannel(new ChannelInitializer<SerialChannel>() {
            @Override
            protected void initChannel(SerialChannel ch) throws Exception {
                ch.pipeline().addLast(new DemoCodec());
            }
        });
        for (int i = 1; i <= 10; i++) {
            DummyHandler.INSTANCE.fireEvent(1);
        }
        channel.close().sync();
    }

   
    private static class DemoCodec extends ByteToMessageCodec<Object> {
        @Override
        protected void encode(ChannelHandlerContext channelHandlerContext, Object o, ByteBuf byteBuf) throws Exception {
            // do nothing here
        }

        @Override
        protected void decode(ChannelHandlerContext channelHandlerContext, ByteBuf byteBuf, List<Object> list) throws Exception {
            byteBuf.markReaderIndex();
            StringBuilder sb = new StringBuilder();
            for (int i = 1; i <= byteBuf.readableBytes(); i++) {
                sb.append(byteBuf.readByte()).append(", ");
            }
            byteBuf.resetReaderIndex();
            logger.debug("We currently have {} readable bytes: {}", byteBuf.readableBytes(), sb.toString());
        }
    }
}