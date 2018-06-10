import os
from flask import Flask
from flask import request
from flask_json import as_json
from flask_json import FlaskJSON


from .exceptions import ConfigurationException
from .webhooker import WebHooker


class App(Flask):
    def __init__(self, *args, auth_key, routes):
        super().__init__("bots")
        self.config['JSON_ADD_STATUS'] = False
        self.config['JSON_USE_ENCODE_METHODS'] = True
        json = FlaskJSON(app=self)

        self._validate_routes(routes)

        @self.route('/<path:path>', methods=["POST"])
        @as_json
        def catch_all(path):
            auth = request.headers.get("Authorization")
            if auth is None or auth != auth_key:
                return dict(), 401

            bot = routes.get(path)
            if bot is None:
                return NotFound()

            return self._handler(
                bot_class=bot,
                request=request.get_json(force=True)
            )

    def _validate_routes(self, routes):
        for route, bot_class in routes.items():
            if not issubclass(bot_class, WebHooker):
                raise ConfigurationException(
                    route + " has invalid class " + type(bot_class))

    def _handler(self, *args, bot_class, request):
        method_name = request.get("current_block", "__not_found")
        bot = bot_class()
        resp = getattr(bot, method_name, NotFound)(**request)
        if not resp:
            resp = (bot, 200)
        return resp


def NotFound(*args, **kwargs):
    return dict(), 404
