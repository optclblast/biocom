-module(hearsayd_app).
-behaviour(application).
-export([start/2, stop/1]).

% todo define and include logger.hrl

start(_StartType, _StartArgs) -> 
        {ok, Dispatch} = application:get_env(hearsay_d, Key).

stop(_State) ->
        ok.