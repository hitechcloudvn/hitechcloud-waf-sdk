"""Base resource class."""


class BaseResource:
    def __init__(self, http):
        self._http = http
