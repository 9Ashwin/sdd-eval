from __future__ import annotations

from dataclasses import dataclass
from threading import Lock


@dataclass
class Task:
    id: int
    title: str
    done: bool = False
    priority: str = "Medium"


class Store:
    def __init__(self):
        self._mu = Lock()
        self._tasks: list[Task] = []
        self._next_id = 1

    def list(self) -> list[Task]:
        with self._mu:
            return list(self._tasks)

    def list_by_priority(self, priority: str) -> list[Task]:
        with self._mu:
            return [t for t in self._tasks if t.priority == priority]

    def create(self, title: str, priority: str = "Medium") -> Task:
        with self._mu:
            t = Task(id=self._next_id, title=title, priority=priority or "Medium")
            self._next_id += 1
            self._tasks.append(t)
            return t

    def delete(self, id: int) -> bool:
        with self._mu:
            for i, t in enumerate(self._tasks):
                if t.id == id:
                    self._tasks.pop(i)
                    return True
            return False

    def toggle_done(self, id: int, done: bool) -> Task | None:
        with self._mu:
            for t in self._tasks:
                if t.id == id:
                    t.done = done
                    return t
            return None


VALID_PRIORITIES = {"Low", "Medium", "High"}


def is_valid_priority(p: str) -> bool:
    return p in VALID_PRIORITIES
