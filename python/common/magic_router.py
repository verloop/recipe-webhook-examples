from importlib import import_module
import pkgutil


def magic_router(*args, module):
    module_name = module.__name__
    importer = pkgutil.get_importer(module.__path__[0])
    iter = pkgutil.iter_importer_modules(importer)
    return {name: bot(module_name, name) for name, _ in iter}


def bot(module, name):
    print(module, name)
    mod = import_module(name="{}.{}".format(module, name))
    bot_class = getattr(mod, "Bot", None)
    if bot_class is None:
        raise ImportError("{} has no class named Bot")

    return bot_class
