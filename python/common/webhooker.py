
class WebHooker(object):

    def __init__(self):
        self.state = {}
        self._variables = {}
        self._exports = {}
        self._next_block = ""

    def __json__(self):
        return dict(
            state=self.state,
            variables=self._variables,
            exports=self._exports,
            next_block=self._next_block,
        )

    def next_block(self, *args, name):
        self._next_block = name

    def export(self, **variables):
        self._exports.update(variables)

    def variable(self, **variables):
        self._variables.update(variables)

    def __not_found(self, *args, **kwargs):
        return dict(), 404
