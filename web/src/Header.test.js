import { render, screen, waitFor } from "@testing-library/react";
import { MockedProvider } from "@apollo/client/testing";
import Header, { SERVER_INFO } from "./Header";

const httpMocks = [
  {
    request: {
      query: SERVER_INFO,
    },
    result: {
      data: {
        serverInfo: {
          build_info: {
            version: "foo-build",
          },
          request_address: "foo-request-address",
          request_port: "foo-request-port",
          web_port: "foo-web-port",
          protocol: "http",
        },
      },
    },
  },
];

const wsMocks = [
  {
    request: {
      query: SERVER_INFO,
    },
    result: {
      data: {
        serverInfo: {
          build_info: {
            version: "foo-build",
          },
          request_address: "foo-request-address",
          request_port: "foo-request-port",
          web_port: "foo-web-port",
          protocol: "ws",
        },
      },
    },
  },
];

describe("Header", () => {
  test("has title", () => {
    render(
      <MockedProvider mocks={httpMocks} addTypename={false}>
        <Header />
      </MockedProvider>
    );

    const title = screen.getByText(/Request Hole/i);
    expect(title).toBeInTheDocument();
  });

  test("has send request", () => {
    render(
      <MockedProvider mocks={httpMocks} addTypename={false}>
        <Header />
      </MockedProvider>
    );

    const sendRequest = screen.getByText(/Send a Request/i);
    expect(sendRequest).toBeInTheDocument();
  });

  test("has view project on github", () => {
    render(
      <MockedProvider mocks={httpMocks} addTypename={false}>
        <Header />
      </MockedProvider>
    );

    const viewProjectOnGithub = screen.getByText(/View Project on Github/i);
    expect(viewProjectOnGithub).toBeInTheDocument();
  });
});

describe("Header fetches and renders server info", () => {
  test("shows initial loading", () => {
    render(
      <MockedProvider mocks={httpMocks} addTypename={false}>
        <Header />
      </MockedProvider>
    );

    const loading = screen.getByText(/Loading server info.../i);
    expect(loading).toBeInTheDocument();
  });

  test("shows error if failed", async () => {
    const errorMock = [
      {
        request: {
          query: SERVER_INFO,
        },
        error: new Error("An error occurred"),
      },
    ];
    render(
      <MockedProvider mocks={errorMock} addTypename={false}>
        <Header />
      </MockedProvider>
    );

    await waitFor(() => new Promise((resolve) => setTimeout(resolve, 0)));

    const loading = screen.queryByText(/Loading server info.../i);
    const err = screen.queryByText(/Failed to load server info./i);
    expect(loading).not.toBeInTheDocument();
    expect(err).toBeInTheDocument();
  });

  test("has version", async () => {
    render(
      <MockedProvider mocks={httpMocks} addTypename={false}>
        <Header />
      </MockedProvider>
    );

    await waitFor(() => new Promise((resolve) => setTimeout(resolve, 0)));

    const buildInfo = screen.getByText(/foo-build/i);
    expect(buildInfo).toBeInTheDocument();
  });

  test("has http server address", async () => {
    render(
      <MockedProvider mocks={httpMocks} addTypename={false}>
        <Header />
      </MockedProvider>
    );

    await waitFor(() => new Promise((resolve) => setTimeout(resolve, 0)));

    const buildInfo = screen.getByText(
      /Listening on: http:\/\/foo-request-address:foo-request-port/i
    );
    expect(buildInfo).toBeInTheDocument();
  });

  test("has ws server address", async () => {
    render(
      <MockedProvider mocks={wsMocks} addTypename={false}>
        <Header />
      </MockedProvider>
    );

    await waitFor(() => new Promise((resolve) => setTimeout(resolve, 0)));

    const buildInfo = screen.getByText(
      /Listening on: ws:\/\/foo-request-address:foo-request-port/i
    );
    expect(buildInfo).toBeInTheDocument();
  });
});
