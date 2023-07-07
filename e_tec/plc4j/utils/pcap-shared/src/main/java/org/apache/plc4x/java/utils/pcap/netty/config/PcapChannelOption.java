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
package org.apache.plc4x.java.utils.pcap.netty.config;

import io.netty.channel.ChannelOption;

public class PcapChannelOption {

    protected PcapChannelOption(){
        throw new IllegalAccessError("Classe utile!");
    }

    public static final ChannelOption<Object> SUPPORT_VLANS =
        ChannelOption.valueOf(Integer.class, "SUPPORT_VLANS");

    /**
     * Option to restrict the captures based on packet port.
     */
    public static final ChannelOption<Object> PORT =
        ChannelOption.valueOf(Integer.class, "PORT");

    /**
     * Option to restrict the captures based on TCP protocol ids.
     */
    public static final ChannelOption<Object> PROTOCOL_ID =
        ChannelOption.valueOf(Integer.class, "PROTOCOL_ID");

    /**
     * Option for providing a PacketHandler, that intercepts the captured packets
     * before passing the data into the channel.
     */
    public static final ChannelOption<Object> PACKET_HANDLER =
        ChannelOption.valueOf(Integer.class, "PACKET_HANDLER");

    /**
     * Option for automatically resolving the remote ips MAC address with an
     * ARP request.
     */
    public static final ChannelOption<Object> RESOLVE_MAC_ADDRESS =
        ChannelOption.valueOf("RESOLVE_MAC_ADDRESS");


}
