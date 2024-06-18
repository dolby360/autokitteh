class AttrDict(dict):
    """Allow attribute access to dictionary keys.

    >>> config = AttrDict({'server': {'port': 8080}, 'debug': True})
    >>> config.server.port
    8080
    >>> config.debug
    True
    """
    def __getattr__(self, name):
        try:
            value = self[name]
            if isinstance(value, dict):
                value = AttrDict(value)
            return value
        except KeyError:
            raise AttributeError(name)

    def __setattr__(self, attr, value):
        # The default __getattr__ doesn't fail but also don't change values
        cls = self.__class__.__name__
        raise NotImplementedError(f'{cls} does not support setting attributes')

