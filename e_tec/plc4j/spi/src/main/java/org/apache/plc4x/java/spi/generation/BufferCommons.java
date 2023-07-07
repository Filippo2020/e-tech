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
package org.apache.plc4x.java.spi.generation;

import org.apache.commons.lang3.StringUtils;

import java.util.Optional;
import java.util.stream.Stream;

public interface BufferCommons {
    String RWDATATYPEKEY = "dataType";
    String RWBITLENGTHKEY = "bitLength";
    String RWSTRINGREPRESENTATIONKEY = "stringRepresentation";
    String RWBITKEY = "bit";
    String RWBYTEKEY = "byte";
    String RWUINTKEY = "uint";
    String RWINTKEY = "int";
    String RWFLOATKEY = "float";
    String RWSTRINGKEY = "string";
    String RWENCODINGKEY = "encoding";
    String RWISLISTKEY = "isList";

    default String sanitizeLogicalName(String logicalName) {
        if (StringUtils.isBlank(logicalName)) {
            return "value";
        }
        return logicalName;
    }

    default boolean isToBeRenderedAsList(WithReaderArgs... readerArgs) {
        return isToBeRenderedAsList(Stream.of(readerArgs).map(WithReaderWriterArgs.class::cast).toArray(WithReaderWriterArgs[]::new));
    }

    default boolean isToBeRenderedAsList(WithWriterArgs... writerArgs) {
        return isToBeRenderedAsList(Stream.of(writerArgs).map(WithReaderWriterArgs.class::cast).toArray(WithReaderWriterArgs[]::new));
    }

    default boolean isToBeRenderedAsList(WithReaderWriterArgs... readerWriterArgs) {
        for (WithReaderWriterArgs arg : readerWriterArgs) {
            if (arg instanceof WithRenderAsList) {
                return ((WithRenderAsList) arg).renderAsList();
            }
        }
        return false;
    }

    default Optional<String> extractAdditionalStringRepresentation(WithReaderArgs... readerArgs) {
        return extractAdditionalStringRepresentation(Stream.of(readerArgs).map(WithReaderWriterArgs.class::cast).toArray(WithReaderWriterArgs[]::new));
    }

    default Optional<String> extractAdditionalStringRepresentation(WithWriterArgs... writerArgs) {
        return extractAdditionalStringRepresentation(Stream.of(writerArgs).map(WithReaderWriterArgs.class::cast).toArray(WithReaderWriterArgs[]::new));
    }

    default Optional<String> extractAdditionalStringRepresentation(WithReaderWriterArgs... readerWriterArgs) {
        for (WithReaderWriterArgs arg : readerWriterArgs) {
            if (arg instanceof WithAdditionalStringRepresentation) {
                return Optional.of(((WithAdditionalStringRepresentation) arg).stringRepresentation());
            }
        }
        return Optional.empty();
    }
}
