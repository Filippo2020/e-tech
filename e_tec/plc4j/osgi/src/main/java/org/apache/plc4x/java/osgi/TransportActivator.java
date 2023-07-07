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
package org.apache.plc4x.java.osgi;

import java.util.ArrayList;
import java.util.List;
import java.util.Properties;

import org.osgi.framework.BundleActivator;
import org.osgi.framework.BundleContext;
import org.osgi.framework.ServiceRegistration;
import org.osgi.framework.wiring.BundleWiring;
import org.apache.plc4x.java.spi.transport.Transport;

import java.util.ServiceLoader;

public class TransportActivator implements BundleActivator {

    private final List<ServiceRegistration<Transport>> registrations = new ArrayList<>();
    private static final String TRASPORTCODE ="org.apache.plc4x.transport.code";
    private static final String TRANSPORTNAME ="org.apache.plc4x.transport.name";


    @Override
    public void start(BundleContext context) throws Exception {
        ServiceLoader<Transport> transports = ServiceLoader.load(Transport.class, context.getBundle().adapt(BundleWiring.class).getClassLoader());
        for (Transport transport : transports) {
            Properties props = new Properties();
            props.put(TRASPORTCODE, transport.getTransportCode());
            props.put(TRANSPORTNAME, transport.getTransportName());
        }
    }

    @Override
    public void stop(BundleContext context) {
        registrations.forEach(ServiceRegistration::unregister);
        registrations.clear();
    }
}




