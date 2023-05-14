"""Please keep the code clean!"""

from threading import Thread
from typing import Any, Callable, Dict, Optional


class GameObject(object):
    ...


class Player(GameObject):
    ...


class Blob(GameObject):
    ...


class Dot(GameObject):
    ...


class Ejection(GameObject):
    ...


class Hole(GameObject):
    ...


class Item(GameObject):
    ...


class Spell(GameObject):
    ...


class World(GameObject):
    ...


class V2API(Thread):
    def __init__(self) -> None:
        super().__init__(name="Nebul0us Python API", daemon=True)
        self.apiMethods: Dict[str, Callable[..., Any]] = {}

    def register(self, name: Optional[str] = None) -> Callable[..., Any]:
        def decorator(func: Callable[..., Any]) -> Callable[..., Any]:
            self.apiMethods[name or func.__name__] = func
            return func

        return decorator

    def connect(self) -> None:
        raise NotImplementedError("plugin API not implemented")

    def run(self) -> None:
        raise NotImplementedError("plugin API not implemented")
