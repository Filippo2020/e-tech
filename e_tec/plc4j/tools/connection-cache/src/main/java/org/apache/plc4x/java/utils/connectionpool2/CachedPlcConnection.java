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
package org.apache.plc4x.java.utils.connectionpool2;

import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.exceptions.PlcConnectionException;
import org.apache.plc4x.java.api.exceptions.PlcRuntimeException;
import org.apache.plc4x.java.api.exceptions.PlcUnsupportedOperationException;
import org.apache.plc4x.java.api.messages.*;
import org.apache.plc4x.java.api.metadata.PlcConnectionMetadata;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.concurrent.CompletableFuture;
import java.util.concurrent.Executors;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.TimeoutException;


/**
 * Wrapper around a PlcConnection which interacts with the {@link CachedDriverManager}.
 */
public class CachedPlcConnection implements PlcConnection, PlcConnectionMetadata {
    String frase = "Request finished with exception. Reporting Connection as Broken";

    private static final Logger logger = LoggerFactory.getLogger(CachedPlcConnection.class);
    private static final String S1="Unnecessary to cancel the request!";
    private static final String S2="Trying to executing Request {}";
    private static final String S3="Trying to execute a Request on a closed Connection!";
    private static final String S4="Trying to build a Request on a closed Connection!";
    private static final String S5="Timing out the PLC request!";
    private static final String S6="Response did not finish in Time!";
    private static final String S7="Request finsihed successfull!";
    private static final String S8="Request failed";
    private static final String S9="Executing Request {}";
    private static final ScheduledExecutorService schedulerExecutor =
        Executors.newScheduledThreadPool(10);

    private final CachedDriverManager parent;
    private PlcConnection activeConnection;
    private volatile boolean closed = false;

    public CachedPlcConnection(CachedDriverManager parent, PlcConnection activeConnection) {
        this.parent = parent;
        this.activeConnection = activeConnection;
    }

    @Override
    public void connect() throws PlcConnectionException {
        // Do nothing
        logger.warn(".connect() is called on a Cached Connection. This has no effect but should not happen.");
    }

    @Override
    public boolean isConnected() {
        if (closed) {
            return false;
        } else {
            return activeConnection.isConnected();
        }
    }

    private CompletableFuture<? extends PlcBrowseResponse> wrapBrowseWithTimeout(CompletableFuture<? extends PlcBrowseResponse> future, long timeoutMillis) {
        //schedule watcher
        final CompletableFuture<PlcBrowseResponse> responseFuture = new CompletableFuture<>();
        schedulerExecutor.schedule(() -> {
            if (!future.isDone()) {
                logger.debug(S5);
                future.cancel(true);
                responseFuture.completeExceptionally(new TimeoutException(S6));
            } else {
                logger.trace(S1);
            }
        }, timeoutMillis, TimeUnit.MILLISECONDS);
        
        future.handle((PlcBrowseResponse plcBrowseResponse, Throwable throwable) -> {
            if (plcBrowseResponse != null) {
                logger.debug(S7);
                responseFuture.complete(plcBrowseResponse);
            } else {
                logger.debug(S8, throwable);
                responseFuture.completeExceptionally(throwable);
            }
            return null;
        });

        return responseFuture;
    }

    private CompletableFuture<? extends PlcReadResponse> wrapReadWithTimeout(CompletableFuture<? extends PlcReadResponse> future, long timeoutMillis) {
        //schedule watcher
        final CompletableFuture<PlcReadResponse> responseFuture = new CompletableFuture<>();
        schedulerExecutor.schedule(() -> {
            if (!future.isDone()) {
                logger.debug(S5);
                future.cancel(true);
                responseFuture.completeExceptionally(new TimeoutException(S6));
            } else {
                logger.trace(S1);
            }
        }, timeoutMillis, TimeUnit.MILLISECONDS);
        
        future.handle((PlcReadResponse plcReadResponse, Throwable throwable) -> {
            if (plcReadResponse != null) {
                logger.debug(S7);
                responseFuture.complete(plcReadResponse);
            } else {
                logger.debug(S8, throwable);
                responseFuture.completeExceptionally(throwable);
            }
            return null;
        });

        return responseFuture;
    }

    private CompletableFuture<? extends PlcWriteResponse> wrapWriteWithTimeout(CompletableFuture<? extends PlcWriteResponse> future, long timeoutMillis) {
        //schedule watcher
        final CompletableFuture<PlcWriteResponse> responseFuture = new CompletableFuture<>();
        schedulerExecutor.schedule(() -> {
            if (!future.isDone()) {
                logger.debug(S5);
                future.cancel(true);
                responseFuture.completeExceptionally(new TimeoutException(S6));
            } else {
                logger.trace(S1);
            }
        }, timeoutMillis, TimeUnit.MILLISECONDS);
        
        future.handle((PlcWriteResponse plcWriteResponse, Throwable throwable) -> {
            if (plcWriteResponse != null) {
                logger.debug(S7);
                responseFuture.complete(plcWriteResponse);
            } else {
                logger.debug(S8, throwable);
                responseFuture.completeExceptionally(throwable);
            }
            return null;
        });

        return responseFuture;
    }

    public CompletableFuture<? extends PlcBrowseResponse> execute(PlcBrowseRequest request) {
        logger.trace(S2, request);
        if (closed) {
            throw new IllegalStateException(S3);
        }
        try {
            logger.trace(S9, request);
            final CompletableFuture<? extends PlcBrowseResponse> responseFuture = wrapBrowseWithTimeout(request.execute(), 5_000);
            // The following code handles the case, that a read fails (which is handled async and thus not really connected
            // to the connection, yet
            // Thus, we register our own listener who gets the notification and reports the connection as broken
            
            return responseFuture.handleAsync((PlcBrowseResponse plcBrowseResponse, Throwable throwable) -> {
                if (throwable != null) {
                    // Do something here...
                    logger.warn(frase, throwable);
                    closeConnectionExceptionally(null);
                }
                return plcBrowseResponse;
            });

        } catch (Exception e) {
            return (CompletableFuture<? extends PlcBrowseResponse>) closeConnectionExceptionally(e);
        }
    }

    public CompletableFuture<? extends PlcBrowseResponse> executeWithInterceptor(PlcBrowseRequest request, PlcBrowseRequestInterceptor interceptor) {
        logger.trace("Trying to executing Request with interceptor {}", request);
        if (closed) {
            throw new IllegalStateException(S3);
        }
        try {
            logger.trace(S9, request);
            final CompletableFuture<? extends PlcBrowseResponse> responseFuture = wrapBrowseWithTimeout(request.executeWithInterceptor(interceptor), 5_000);
            // The following code handles the case, that a read fails (which is handled async and thus not really connected
            // to the connection, yet
            // Thus, we register our own listener who gets the notification and reports the connection as broken
            
            return responseFuture.handleAsync((PlcBrowseResponse plcBrowseResponse, Throwable throwable) -> {
                if (throwable != null) {
                    // Do something here...
                    logger.warn(frase, throwable);
                    closeConnectionExceptionally(null);
                }
                return plcBrowseResponse;
            });

        } catch (Exception e) {
            return (CompletableFuture<? extends PlcBrowseResponse>) closeConnectionExceptionally(e);
        }
    }

    /**
     * Executes the Request.
     */
    public CompletableFuture<? extends PlcReadResponse> execute(PlcReadRequest request) {
        logger.trace(S2, request);
        if (closed) {
            throw new IllegalStateException(S3);
        }
        try {
            logger.trace(S9, request);
             final CompletableFuture<? extends PlcReadResponse> responseFuture = wrapReadWithTimeout(request.execute(), 5_000);

            // The following code handles the case, that a read fails (which is handled async and thus not really connected
            // to the connection, yet
            // Thus, we register our own listener who gets the notification and reports the connection as broken
            
            return responseFuture.handleAsync((PlcReadResponse plcReadResponse, Throwable throwable) -> {
                if (throwable != null) {
                    // Do something here...
                    logger.warn(frase, throwable);
                    closeConnectionExceptionally(null);
                }
                return plcReadResponse;
            });

        } catch (Exception e) {
            return (CompletableFuture<? extends PlcReadResponse>) closeConnectionExceptionally(e);
        }
    }

    public CompletableFuture<? extends PlcWriteResponse> execute(PlcWriteRequest request) {
        logger.trace(S2, request);
        if (closed) {
            throw new IllegalStateException(S3);
        }
        try {
            logger.trace(S9, request);
            final CompletableFuture<? extends PlcWriteResponse> responseFuture = wrapWriteWithTimeout(request.execute(), 5_000);
            // The following code handles the case, that a read fails (which is handled async and thus not really connected
            // to the connection, yet
            // Thus, we register our own listener who gets the notification and reports the connection as broken
            
            return responseFuture.handleAsync((PlcWriteResponse plcWriteResponse, Throwable throwable) -> {
                if (throwable != null) {
                    // Do something here...
                    logger.warn(frase, throwable);
                    closeConnectionExceptionally(null);
                }
                return plcWriteResponse;
            });

        } catch (Exception e) {
            return (CompletableFuture<? extends PlcWriteResponse>) closeConnectionExceptionally(e);
        }
    }

    private CompletableFuture<? extends PlcResponse> closeConnectionExceptionally(Exception e) {
        // First, close this connection and allow no further operations on it!
        this.closed = true;
        // Return the Connection as invalid
        parent.handleBrokenConnection();
        // Invalidate Connection
        this.activeConnection = null;
        // Throw Exception
        throw new PlcRuntimeException("Unable to finish Request!", e);
    }

    PlcConnection getActiveConnection() {
        return this.activeConnection;
    }

    @Override
    public synchronized void close() throws Exception {
        logger.debug("Closing cached connection and returning borrowed connection to pool.");
        // First, close this connection and allow no further operations on it!
        this.closed = true;
        // Return the Connection
        parent.returnConnection();
        // Invalidate Connection
        this.activeConnection = null;
    }

    @Override
    public PlcConnectionMetadata getMetadata() {
        if (closed) {
            throw new IllegalStateException("Trying to get Metadata on a closed Connection!");
        } else {
            return this;
        }
    }

    @Override
    public CompletableFuture<Void> ping() {
        CompletableFuture<Void> future = new CompletableFuture<>();
        future.completeExceptionally(new PlcUnsupportedOperationException("The connection does not support pinging"));
        return future;
    }

    @Override
    public PlcBrowseRequest.Builder browseRequestBuilder() {
        if (closed) {
            throw new IllegalStateException(S4);
        }
        return new CachedBrowseRequestBuilder(this, this.getActiveConnection().browseRequestBuilder());
    }

    @Override
    public PlcReadRequest.Builder readRequestBuilder() {
        if (closed) {
            throw new IllegalStateException(S4);
        }
        return new CachedReadRequestBuilder(this, this.getActiveConnection().readRequestBuilder());
    }

    @Override
    public PlcWriteRequest.Builder writeRequestBuilder() {
        if (closed) {
            throw new IllegalStateException(S4);
        }
        return new CachedWriteRequestBuilder(this, this.getActiveConnection().writeRequestBuilder());
    }

    @Override
    public PlcSubscriptionRequest.Builder subscriptionRequestBuilder() {
        throw new UnsupportedOperationException();
    }

    @Override
    public PlcUnsubscriptionRequest.Builder unsubscriptionRequestBuilder() {
        throw new UnsupportedOperationException();
    }

    @Override
    public boolean canBrowse() {
        if (closed) {
            return false;
        } else {
            return this.activeConnection.getMetadata().canBrowse();
        }
    }

    @Override
    public boolean canRead() {
        if (closed) {
            return false;
        } else {
            return this.activeConnection.getMetadata().canRead();
        }
    }

    @Override
    public boolean canWrite() {
        if (closed) {
            return false;
        } else {
            return this.activeConnection.getMetadata().canWrite();
        }
    }

    @Override
    public boolean canSubscribe() {
        if (closed) {
            return false;
        } else {
            return this.activeConnection.getMetadata().canSubscribe();
        }
    }
}

