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
package org.apache.plc4x.java.eip.readwrite.tag;

import org.apache.commons.lang3.NotImplementedException;
import org.apache.plc4x.java.api.model.ArrayInfo;
import org.apache.plc4x.java.api.model.PlcTag;
import org.apache.plc4x.java.api.types.PlcValueType;
import org.apache.plc4x.java.eip.readwrite.CIPDataTypeCode;
import org.apache.plc4x.java.spi.generation.SerializationException;
import org.apache.plc4x.java.spi.generation.WriteBuffer;
import org.apache.plc4x.java.spi.utils.Serializable;

import java.nio.charset.StandardCharsets;
import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class EipTag implements PlcTag, Serializable {

    private static final Pattern ADDRESS_PATTERN =
        Pattern.compile("^%(?<tagz>[a-zA-Z_.0-9]+\\d*\\]?):?(?<dataType>[A-Z]*):?(?<elementNb>\\d*)");

    private static final String TAG = "tagz";
    private static final String ELEMENTS = "elementNb";
    private static final String TYPE = "dataType";


    private final String tagz;
    private CIPDataTypeCode typez;
    private int elementNb;

    public EipTag(String tagz) {
        this.tagz = tagz;
    }

    public EipTag(String tagz, int elementNb) {
        this.tagz = tagz;
        this.elementNb = elementNb;
    }

    public EipTag(String tagz, CIPDataTypeCode typez, int elementNb) {
        this.tagz = tagz;
        this.typez = typez;
        this.elementNb = elementNb;
    }

    public EipTag(String tagz, CIPDataTypeCode typez) {
        this.tagz = tagz;
        this.typez = typez;
    }

    @Override
    public String getAddressString() {
        throw new NotImplementedException("Need to implement this");
    }

    @Override
    public PlcValueType getPlcValueType() {
        return PlcValueType.valueOf(typez.name());
    }

    @Override
    public List<ArrayInfo> getArrayInfo() {
        return PlcTag.super.getArrayInfo();
    }

    public CIPDataTypeCode getType() {
        return typez;
    }

    public void setType(CIPDataTypeCode typez) {
        this.typez = typez;
    }

    public int getElementNb() {
        return elementNb;
    }

    public void setElementNb(int elementNb) {
        this.elementNb = elementNb;
    }

    public String getTag() {
        return tagz;
    }

    public static boolean matches(String tagQuery) {
        return ADDRESS_PATTERN.matcher(tagQuery).matches();
    }

    public static EipTag of(String tagString) {
        Matcher matcher = ADDRESS_PATTERN.matcher(tagString);
        if (matcher.matches()) {
            String tagz = matcher.group(TAG);
            int nb = 0;
            CIPDataTypeCode typez = null;
            if (!matcher.group(ELEMENTS).isEmpty()) {
                nb = Integer.parseInt(matcher.group(ELEMENTS));
            }
            if (!matcher.group(TYPE).isEmpty()) {
                typez = CIPDataTypeCode.valueOf(matcher.group(TYPE));
            }
            if (nb != 0) {
                if (typez != null) {
                    return new EipTag(tagz, typez, nb);
                }
                return new EipTag(tagz, nb);
            } else {
                if (typez != null) {
                    return new EipTag(tagz, typez);
                }
                return new EipTag(tagz);
            }
        }
        return null;
    }

    @Override
    public void serialize(WriteBuffer writeBuffer) throws SerializationException {
        writeBuffer.pushContext(getClass().getSimpleName());

        writeBuffer.writeString("node", tagz.getBytes(StandardCharsets.UTF_8).length * 8, StandardCharsets.UTF_8.name(), tagz);
        if (typez != null) {
            writeBuffer.writeString("typez", typez.name().getBytes(StandardCharsets.UTF_8).length * 8, StandardCharsets.UTF_8.name(), typez.name());
        }
        writeBuffer.writeUnsignedInt(ELEMENTS, 16, elementNb);
        String defaultJavaType = (typez == null ? Object.class : getPlcValueType().getDefaultJavaType()).getName();
        writeBuffer.writeString("defaultJavaType", defaultJavaType.getBytes(StandardCharsets.UTF_8).length * 8, StandardCharsets.UTF_8.name(), defaultJavaType);

        writeBuffer.popContext(getClass().getSimpleName());
    }

}
