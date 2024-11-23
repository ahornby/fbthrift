/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

package test.fixtures.interactions;

import com.facebook.thrift.client.*;
import java.util.*;

public class BoxServiceBlockingReactiveWrapper 
    implements BoxService.Reactive {
    private final BoxService _delegate;

    public BoxServiceBlockingReactiveWrapper(BoxService _delegate) {
        
        this._delegate = _delegate;
    }

    @java.lang.Override
    public void dispose() {
        _delegate.close();
    }

    @java.lang.Override
    public reactor.core.publisher.Mono<test.fixtures.interactions.ShouldBeBoxed> getABoxSession(final test.fixtures.interactions.ShouldBeBoxed req) {
        reactor.core.publisher.Mono<test.fixtures.interactions.ShouldBeBoxed> _m = reactor.core.publisher.Mono.create(_sink -> {
            try {
                reactor.util.context.ContextView _contextView = _sink.contextView();
                com.facebook.nifty.core.RequestContext
                    .tryContextView(_contextView)
                    .ifPresent(com.facebook.nifty.core.RequestContexts::setCurrentContext);
                _sink.success(_delegate.getABoxSession(req));
            } catch (Throwable _e) {
                _sink.error(_e);
            }
        });

        if (!com.facebook.thrift.util.resources.RpcResources.isForceExecutionOffEventLoop()) {
            _m = _m.subscribeOn(com.facebook.thrift.util.resources.RpcResources.getOffLoopScheduler());
        }

        return _m;
    }

    public class BoxedInteractionImpl implements BoxedInteraction {
        private BoxService.BoxedInteraction _delegateInteraction;

        BoxedInteractionImpl(BoxService.BoxedInteraction delegateInteraction) {
            this._delegateInteraction = delegateInteraction;
        }

        @java.lang.Override
        public reactor.core.publisher.Mono<test.fixtures.interactions.ShouldBeBoxed> getABox() {
                reactor.core.publisher.Mono<test.fixtures.interactions.ShouldBeBoxed> _m = reactor.core.publisher.Mono.create(_sink -> {
                    try {
                        reactor.util.context.ContextView _contextView = _sink.contextView();
                        com.facebook.nifty.core.RequestContext
                            .tryContextView(_contextView)
                            .ifPresent(com.facebook.nifty.core.RequestContexts::setCurrentContext);
                        _sink.success(_delegateInteraction.getABox());
                    } catch (Throwable _e) {
                        _sink.error(_e);
                    }
                });

                if (!com.facebook.thrift.util.resources.RpcResources.isForceExecutionOffEventLoop()) {
                    _m = _m.subscribeOn(com.facebook.thrift.util.resources.RpcResources.getOffLoopScheduler());
                }

                return _m;
        }

        @java.lang.Override
        public reactor.core.publisher.Mono<test.fixtures.interactions.ShouldBeBoxed> getABox(RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }

        @java.lang.Override
        public reactor.core.publisher.Mono<ResponseWrapper<test.fixtures.interactions.ShouldBeBoxed>> getABoxWrapper(RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }

        @java.lang.Override
        public void dispose() {}
    }

    public BoxedInteraction createBoxedInteraction() {
        return new BoxedInteractionImpl(_delegate.createBoxedInteraction());
    }
}
