from fastapi import FastAPI, HTTPException
from pydantic import BaseModel

from store import Store, is_valid_priority

app = FastAPI()
store = Store()


class CreateRequest(BaseModel):
    title: str
    priority: str = ""


class ToggleRequest(BaseModel):
    done: bool


@app.get("/tasks")
def list_tasks(priority: str | None = None):
    if priority:
        return store.list_by_priority(priority)
    return store.list()


@app.post("/tasks", status_code=201)
def create_task(req: CreateRequest):
    priority = req.priority or "Medium"
    if not is_valid_priority(priority):
        raise HTTPException(400, detail="invalid priority")
    return store.create(req.title, priority)


@app.delete("/tasks/{task_id}", status_code=204)
def delete_task(task_id: int):
    if not store.delete(task_id):
        raise HTTPException(404, detail="not found")


@app.patch("/tasks/{task_id}")
def toggle_task(task_id: int, req: ToggleRequest):
    task = store.toggle_done(task_id, req.done)
    if task is None:
        raise HTTPException(404, detail="not found")
    return task


if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, port=3000)
