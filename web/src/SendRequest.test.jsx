import { expect, describe, test } from "vitest";
import { render, screen, waitFor } from "@testing-library/react";
import "@testing-library/jest-dom";
import { MockedProvider } from "@apollo/client/testing";
import SendRequest, { SERVER_INFO } from "./SendRequest";

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
        },
      },
    },
  },
];

describe("SendRequest", () => {
  test("has form fields", () => {
    render(
      <MockedProvider mocks={mocks} addTypename={false}>
        <SendRequest filters={[]} visible={true} />
      </MockedProvider>,
    );

    expect(screen.getByLabelText(/method/i)).toBeInTheDocument();
    expect(screen.getByLabelText(/url/i)).toBeInTheDocument();
    expect(screen.getByLabelText(/body/i)).toBeInTheDocument();
    expect(
      screen.getByRole("button", { name: "Send Request" }),
    ).toBeInTheDocument();
    expect(screen.getByRole("button", { name: "Close" })).toBeInTheDocument();
  });

  test("visibility prop hides component", () => {
    render(
      <MockedProvider mocks={mocks} addTypename={false}>
        <SendRequest filters={[]} visible={true} />
      </MockedProvider>,
    );

    expect(
      screen.getByRole("button", { name: "Send Request" }),
    ).toBeInTheDocument();
    expect(screen.getByRole("button", { name: "Close" })).toBeInTheDocument();
  });

  test("visibility prop shows component", () => {
    render(
      <MockedProvider mocks={mocks} addTypename={false}>
        <SendRequest visible={false} />
      </MockedProvider>,
    );

    expect(screen.queryAllByRole("button")).toHaveLength(0);
  });

  test("filters prop populates method", () => {
    render(
      <MockedProvider mocks={mocks} addTypename={false}>
        <SendRequest filters={["GET", "POST"]} visible={true} />
      </MockedProvider>,
    );

    expect(screen.getByLabelText(/method/i)).toBeInTheDocument();
    expect(screen.getByRole("option", { name: "GET" })).toBeInTheDocument();
    expect(screen.getByRole("option", { name: "POST" })).toBeInTheDocument();
    expect(
      screen.queryByRole("option", { name: "FOOBAR" }),
    ).not.toBeInTheDocument();
  });

  test("body has a default value", () => {
    render(
      <MockedProvider mocks={mocks} addTypename={false}>
        <SendRequest filters={[]} visible={true} />
      </MockedProvider>,
    );

    expect(screen.getByLabelText(/body/i)).toBeInTheDocument();

    const defaultValue = JSON.stringify({ hello: "world" });
    expect(screen.getByRole("textbox", { name: /body/i })).toHaveValue(
      defaultValue,
    );
  });
});

describe("SendRequest fetches serverInfo", () => {
  test("populates url with request_address and request_port", async () => {
    render(
      <MockedProvider mocks={mocks} addTypename={false}>
        <SendRequest filters={[]} visible={true} />
      </MockedProvider>,
    );

    expect(screen.getByLabelText(/url/i)).toBeInTheDocument();

    await waitFor(() => new Promise((resolve) => setTimeout(resolve, 0)));

    const expectedUrl = "http://foo-request-address:foo-request-port";
    expect(screen.getByRole("textbox", { name: /url/i })).toHaveValue(
      expectedUrl,
    );
  });
});
