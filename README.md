# MediChat-Users

MediChat-Users is the Backend component of a Chat Provider application that also includes a SPA [MediChat-UI]() written in React with NodeJS and Fastify.

This application is written in go.

It provides the resources necessary to allow user access to the application. These are the resource provided by this app:

- Authorization with JWT to allow user to register with the Application
- Authentication with JWT to allow registered user to sign in
- User resources
- Channel resources
- Websocket resources to communicate among users
- Swagger to get more information about the available resources.


Future Features:

- Ability to add and delete Channel resources
- Ability to update Channel and User resources
- Fix issue with websocket causing channel pollution in the UI
- Adding a video feature with webrtc.

See [CONTRIBUTING](CONTRIBUTING.md) for more on how to use the application
