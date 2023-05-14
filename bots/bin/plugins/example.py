from v2api import V2API


v2API = V2API()


@v2API.register()
def onPlayerJoin() -> None:
    print("a player joined the server!")


@v2API.register()
def onClientConnected() -> None:
    print("connected!")


def main() -> None:
    v2API.run()  # start the API

    # implement custom logic here
    v2API.connect()

    v2API.join()  # run until the API is stopped


if __name__ == "__main__":
    main()
