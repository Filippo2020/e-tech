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
package org.apache.plc4x.java.scraper.config;

import org.apache.plc4x.java.scraper.config.triggeredscraper.JobConfigurationTriggeredImpl;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * @deprecated Scraper is deprecated please use {@link JobConfigurationTriggeredImpl} instead all functions are supplied as well see java-doc of {@link org.apache.plc4x.java.scraper.triggeredscraper.TriggeredScraperImpl}
 */
@SuppressWarnings("removal")
@Deprecated()
public class JobConfigurationClassicImplBuilder {

    private final ScraperConfigurationClassicImplBuilder parent;
    private final String name;
    private final int scrapeRateMs;

    private final List<String> sources = new ArrayList<>();
    private final Map<String, String> tags = new HashMap<>();

    public JobConfigurationClassicImplBuilder(ScraperConfigurationClassicImplBuilder parent, String name, int scrapeRateMs) {
        this.parent = parent;
        this.name = name;
        this.scrapeRateMs = scrapeRateMs;
    }

    public JobConfigurationClassicImplBuilder source(String alias) {
        this.sources.add(alias);
        return this;
    }

    public JobConfigurationClassicImplBuilder tag(String alias, String tagAddress) {
        this.tags.put(alias, tagAddress);
        return this;
    }

    private JobConfigurationClassicImpl buildInternal() {
        return new JobConfigurationClassicImpl(name,null, scrapeRateMs, sources, tags);
    }

    public ScraperConfigurationClassicImplBuilder build() {
        parent.addJobConfiguration(this.buildInternal());
        return this.parent;
    }
}
