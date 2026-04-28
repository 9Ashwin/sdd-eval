from store import Store, is_valid_priority


class TestList:
    def test_empty(self):
        s = Store()
        assert len(s.list()) == 0


class TestCreate:
    def test_basic(self):
        s = Store()
        task = s.create("buy milk", "Medium")
        assert task.id == 1
        assert task.title == "buy milk"
        assert not task.done
        assert task.priority == "Medium"

    def test_default_priority(self):
        s = Store()
        task = s.create("test", "")
        assert task.priority == "Medium"

    def test_with_priority(self):
        s = Store()
        task = s.create("urgent", "High")
        assert task.priority == "High"


class TestIsValidPriority:
    def test_valid(self):
        assert is_valid_priority("Low")
        assert is_valid_priority("Medium")
        assert is_valid_priority("High")

    def test_invalid(self):
        assert not is_valid_priority("urgent")
        assert not is_valid_priority("")


class TestDelete:
    def test_exists(self):
        s = Store()
        s.create("task", "Medium")
        assert s.delete(1)
        assert not s.delete(1)

    def test_not_found(self):
        s = Store()
        assert not s.delete(99)


class TestToggleDone:
    def test_exists(self):
        s = Store()
        s.create("task", "Medium")

        task = s.toggle_done(1, True)
        assert task is not None
        assert task.done
        assert task.id == 1

        task = s.toggle_done(1, False)
        assert task is not None
        assert not task.done

    def test_not_found(self):
        s = Store()
        assert s.toggle_done(99, True) is None


class TestListByPriority:
    def test_filtering(self):
        s = Store()
        s.create("high task", "High")
        s.create("med task", "Medium")
        s.create("low task", "Low")

        high = s.list_by_priority("High")
        assert len(high) == 1
        assert high[0].title == "high task"

        med = s.list_by_priority("Medium")
        assert len(med) == 1

    def test_no_matches(self):
        s = Store()
        s.create("only medium", "Medium")
        assert len(s.list_by_priority("Low")) == 0
