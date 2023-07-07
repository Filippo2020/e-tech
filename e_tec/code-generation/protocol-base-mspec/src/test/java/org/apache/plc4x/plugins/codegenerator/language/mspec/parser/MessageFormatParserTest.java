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
package org.apache.plc4x.plugins.codegenerator.language.mspec.parser;

import org.apache.plc4x.plugins.codegenerator.language.mspec.model.definitions.DefaultComplexTypeDefinition;
import org.apache.plc4x.plugins.codegenerator.language.mspec.model.definitions.DefaultEnumTypeDefinition;
import org.apache.plc4x.plugins.codegenerator.language.mspec.model.definitions.DefaultTypeDefinition;
import org.apache.plc4x.plugins.codegenerator.language.mspec.model.fields.DefaultSimpleField;
import org.apache.plc4x.plugins.codegenerator.language.mspec.model.fields.DefaultTypedField;
import org.apache.plc4x.plugins.codegenerator.language.mspec.model.fields.DefaultTypedNamedField;
import org.apache.plc4x.plugins.codegenerator.language.mspec.model.references.*;
import org.apache.plc4x.plugins.codegenerator.protocol.TypeContext;
import org.apache.plc4x.plugins.codegenerator.types.enums.EnumValue;
import org.apache.plc4x.plugins.codegenerator.types.references.SimpleTypeReference;
import org.junit.jupiter.api.Test;

import java.util.Map;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.InstanceOfAssertFactories.type;
import static org.junit.jupiter.api.Assertions.assertThrows;

class MessageFormatParserTest {

    MessageFormatParser SUT = new MessageFormatParser();

    @Test
    void parseNull() {
        assertThrows(NullPointerException.class, () -> SUT.parse(null));
    }

    @Test
    void parseSomething() {
        TypeContext parse = SUT.parse(getClass().getResourceAsStream("/mspec.example"));
        assertThat(parse)
            .isNotNull();
    }

    @Test
    void parseSomethingElse() {
        TypeContext parse = SUT.parse(getClass().getResourceAsStream("/mspec.example2"));
        assertThat(parse)
            .isNotNull();
    }

    @Test
    void parseNothingElse() {
        TypeContext typeContext = SUT.parse(getClass().getResourceAsStream("/mspec.example3"));
        assertThat(typeContext)
            .extracting(TypeContext::getUnresolvedTypeReferences)
            .extracting(Map::size)
            .isEqualTo(0);

        assertThat(typeContext)
            .extracting(TypeContext::getTypeDefinitions)
            .satisfies(stringTypeDefinitionMap -> {
                assertThat(stringTypeDefinitionMap)
                    .hasEntrySatisfying("A", typeDefinition ->
                        assertThat(typeDefinition)
                            .asInstanceOf(type(DefaultComplexTypeDefinition.class))
                            .satisfies(defaultComplexTypeDefinition -> {
                                assertThat(defaultComplexTypeDefinition)
                                    .extracting(DefaultTypeDefinition::getName)
                                    .isEqualTo("A");
                                assertThat(defaultComplexTypeDefinition)
                                    .extracting(DefaultComplexTypeDefinition::getFields)
                                    .satisfies(fields -> {
                                        assertThat(fields)
                                            .element(0)
                                            .asInstanceOf(type(DefaultSimpleField.class))
                                            .satisfies(defaultSimpleField -> {
                                                assertThat(defaultSimpleField)
                                                    .extracting(DefaultTypedNamedField::getName)
                                                    .isEqualTo("b");
                                                assertThat(defaultSimpleField)
                                                    .extracting(DefaultTypedField::getType)
                                                    .asInstanceOf(type(DefaultComplexTypeReference.class))
                                                    .extracting(DefaultComplexTypeReference::getName)
                                                    .isEqualTo("B");
                                            });
                                    });
                            }))
                    .hasEntrySatisfying("B", typeDefinition ->
                        assertThat(typeDefinition)
                            .asInstanceOf(type(DefaultComplexTypeDefinition.class))
                            .satisfies(defaultComplexTypeDefinition -> {
                                assertThat(defaultComplexTypeDefinition)
                                    .extracting(DefaultTypeDefinition::getName)
                                    .isEqualTo("B");
                            })
                    )
                    .hasEntrySatisfying("C", typeDefinition ->
                        assertThat(typeDefinition)
                            .asInstanceOf(type(DefaultComplexTypeDefinition.class))
                            .satisfies(defaultComplexTypeDefinition -> {
                                assertThat(defaultComplexTypeDefinition)
                                    .extracting(DefaultTypeDefinition::getName)
                                    .isEqualTo("C");
                                assertThat(defaultComplexTypeDefinition)
                                    .extracting(DefaultComplexTypeDefinition::getFields)
                                    .satisfies(fields -> {
                                        assertThat(fields)
                                            .element(0)
                                            .asInstanceOf(type(DefaultSimpleField.class))
                                            .satisfies(defaultSimpleField -> {
                                                assertThat(defaultSimpleField)
                                                    .extracting(DefaultTypedNamedField::getName)
                                                    .isEqualTo("onlyOneField");
                                                assertThat(defaultSimpleField)
                                                    .extracting(DefaultTypedField::getType)
                                                    .asInstanceOf(type(DefaultBooleanTypeReference.class))
                                                    .extracting(DefaultBooleanTypeReference::getBaseType)
                                                    .isEqualTo(SimpleTypeReference.SimpleBaseType.BIT);
                                            });
                                    });
                            })
                    )
                    .hasEntrySatisfying("D", typeDefinition ->
                        assertThat(typeDefinition)
                            .asInstanceOf(type(DefaultEnumTypeDefinition.class))
                            .satisfies(defaultEnumTypeDefinition -> {
                                assertThat(defaultEnumTypeDefinition)
                                    .extracting(DefaultTypeDefinition::getName)
                                    .isEqualTo("D");
                            })
                    )
                    .hasEntrySatisfying("E", typeDefinition ->
                        assertThat(typeDefinition)
                            .asInstanceOf(type(DefaultComplexTypeDefinition.class))
                            .satisfies(defaultComplexTypeDefinition -> {
                                assertThat(defaultComplexTypeDefinition)
                                    .extracting(DefaultComplexTypeDefinition::getName)
                                    .isEqualTo("E");
                                assertThat(defaultComplexTypeDefinition)
                                    .extracting(DefaultComplexTypeDefinition::getFields)
                                    .satisfies(fields ->
                                        assertThat(fields)
                                            .element(0)
                                            .asInstanceOf(type(DefaultSimpleField.class))
                                            .extracting(DefaultTypedNamedField::getName)
                                            .isEqualTo("eField")
                                    );
                            })
                    )
                    .hasEntrySatisfying("Root", typeDefinition ->
                        assertThat(typeDefinition)
                            .asInstanceOf(type(DefaultComplexTypeDefinition.class))
                            .satisfies(defaultComplexTypeDefinition -> {
                                assertThat(defaultComplexTypeDefinition)
                                    .extracting(DefaultComplexTypeDefinition::getName)
                                    .isEqualTo("Root");
                            })
                    );
            });
    }

}