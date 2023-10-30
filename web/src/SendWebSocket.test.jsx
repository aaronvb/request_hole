import { expect, describe, test } from "vitest";
import { render, screen, waitFor } from "@testing-library/react";
import "@testing-library/jest-dom";
import { MockedProvider } from "@apollo/client/testing";
import SendWebSocket, { SERVER_INFO } from "./SendWebSocket";

const mocks = [
  {
    request: {
      query: SERVER_INFO,
    },
    result: {
      data: {
        serverInfo: {
          request_address: "foo-request-address",
          request_port: "foo-request-port",
          protocol: "ws",
        },
      },
    },
  },
];

describe("SendWebSocket", () => {
  test("has form connection fields", () => {
    render(
      <MockedProvider mocks={mocks} addTypename={false}>
        <SendWebSocket filters={[]} visible={true} />
      </MockedProvider>,
    );

    expect(screen.getByLabelText(/url/i)).toBeInTheDocument();
    expect(screen.queryByLabelText(/body/i)).not.toBeInTheDocument();
    expect(screen.getByRole("button", { name: "Connect" })).toBeInTheDocument();
    expect(screen.getByRole("button", { name: "Close" })).toBeInTheDocument();
  });

  test("visibility prop hides component", () => {
    render(
      <MockedProvider mocks={mocks} addTypename={false}>
        <SendWebSocket filters={[]} visible={true} />
      </MockedProvider>,
    );

    expect(screen.getByRole("button", { name: "Connect" })).toBeInTheDocument();
    expect(screen.getByRole("button", { name: "Close" })).toBeInTheDocument();
  });

  test("visibility prop shows component", () => {
    render(
      <MockedProvider mocks={mocks} addTypename={false}>
        <SendWebSocket visible={false} />
      </MockedProvider>,
    );

    expect(screen.queryAllByRole("button")).toHaveLength(0);
  });

  test("body is hidden when not connected", () => {
    render(
      <MockedProvider mocks={mocks} addTypename={false}>
        <SendWebSocket filters={[]} visible={true} />
      </MockedProvider>,
    );

    expect(screen.queryByLabelText(/body/i)).not.toBeInTheDocument();
  });
});

describe("SendWebSocket fetches serverInfo", () => {
  test("populates url with request_address and request_port", async () => {
    render(
      <MockedProvider mocks={mocks} addTypename={false}>
        <SendWebSocket filters={[]} visible={true} />
      </MockedProvider>,
    );

    expect(screen.getByLabelText(/url/i)).toBeInTheDocument();

    await waitFor(() => new Promise((resolve) => setTimeout(resolve, 0)));

    const expectedUrl = "ws://foo-request-address:foo-request-port";
    expect(screen.getByRole("textbox", { name: /url/i })).toHaveValue(
      expectedUrl,
    );
  });
});
