import { render, screen, waitFor } from "@testing-library/react";
import { MockedProvider } from "@apollo/client/testing";
import Requests, {
  ALL_REQUESTS,
  REQUESTS_SUBSCRIPTION,
  CLEAR_REQUESTS,
} from "./Requests";

const mocks = [
  {
    request: {
      query: ALL_REQUESTS,
    },
    result: {
      data: {
        requests: [
          {
            id: "e0977611-3c2f-4494-9977-1db2352065ef",
            fields: {
              method: "GET",
              url: "/",
            },
            headers: {
              Accept: ["*/*"],
              "Accept-Encoding": ["gzip, deflate"],
              "Accept-Language": ["en-US,en;q=0.5"],
              Connection: ["keep-alive"],
              "Content-Type": ["application/json"],
              Origin: ["http://localhost:3000"],
              Referer: ["http://localhost:3000/"],
            },
            param_fields: {
              form: null,
              query: null,
              json: null,
              json_array: null,
            },
            created_at: "2021-07-09T13:41:27-10:00",
          },
        ],
      },
    },
  },
  {
    request: {
      query: REQUESTS_SUBSCRIPTION,
    },
    result: {},
  },
  {
    request: {
      query: CLEAR_REQUESTS,
    },
    result: {
      data: {
        requests: [],
      },
    },
  },
];

describe("Requests", () => {
  test("has clear request button", () => {
    render(
      <MockedProvider mocks={mocks} addTypename={false}>
        <Requests filters={[]} />
      </MockedProvider>
    );

    expect(
      screen.getByRole("button", { name: "Clear Requests" })
    ).toBeInTheDocument();
  });

  test("has hide details button by default", () => {
    render(
      <MockedProvider mocks={mocks} addTypename={false}>
        <Requests filters={[]} />
      </MockedProvider>
    );

    expect(
      screen.getByRole("button", { name: "Hide Details" })
    ).toBeInTheDocument();
  });

  test("has hide filter button set to all by default", () => {
    render(
      <MockedProvider mocks={mocks} addTypename={false}>
        <Requests filters={[]} />
      </MockedProvider>
    );

    expect(
      screen.getByRole("button", { name: "Filter: ALL" })
    ).toBeInTheDocument();
  });
});

describe("Requests fetches and renders requests", () => {
  test("has count of requests", async () => {
    render(
      <MockedProvider mocks={mocks} addTypename={false}>
        <Requests filters={[]} />
      </MockedProvider>
    );

    expect(
      screen.getByRole("heading", { name: "0 Requests" })
    ).toBeInTheDocument();
    expect(screen.getByText(/Loading requests.../i)).toBeInTheDocument();

    await waitFor(() => new Promise((resolve) => setTimeout(resolve, 0)));

    expect(
      screen.getByRole("heading", { name: "1 Request" })
    ).toBeInTheDocument();
  });

  test("clear requests button clears all requests", async () => {
    render(
      <MockedProvider mocks={mocks} addTypename={false}>
        <Requests filters={[]} />
      </MockedProvider>
    );

    await waitFor(() => new Promise((resolve) => setTimeout(resolve, 0)));

    expect(
      screen.getByRole("heading", { name: "1 Request" })
    ).toBeInTheDocument();

    window.confirm = jest.fn(() => true);
    screen.getByRole("button", { name: "Clear Requests" }).click();
    expect(window.confirm).toBeCalled();

    await waitFor(() => new Promise((resolve) => setTimeout(resolve, 0)));

    expect(
      screen.getByRole("heading", { name: "0 Requests" })
    ).toBeInTheDocument();
  });
});

describe("Details", () => {
  test("clicking hide details hides details", async () => {
    render(
      <MockedProvider mocks={mocks} addTypename={false}>
        <Requests filters={[]} />
      </MockedProvider>
    );

    await waitFor(() => new Promise((resolve) => setTimeout(resolve, 0)));

    expect(
      screen.getByRole("heading", { name: "1 Request" })
    ).toBeInTheDocument();

    expect(screen.getByText(/(\d|no) headers/i)).toBeInTheDocument();
    expect(screen.getByText(/(\d|no) params/i)).toBeInTheDocument();
    screen.getByRole("button", { name: "Hide Details" }).click();

    await waitFor(() => {
      expect(screen.queryByText(/(\d|no) headers/i)).not.toBeInTheDocument();
      expect(screen.queryByText(/(\d|no) params/i)).not.toBeInTheDocument();
    });
  });

  test("clicking show details shows details", async () => {
    render(
      <MockedProvider mocks={mocks} addTypename={false}>
        <Requests filters={[]} />
      </MockedProvider>
    );

    screen.getByRole("button", { name: "Hide Details" }).click();

    await waitFor(() => {
      expect(screen.queryByText(/(\d|no) headers/i)).not.toBeInTheDocument();
      expect(screen.queryByText(/(\d|no) params/i)).not.toBeInTheDocument();
    });

    screen.getByRole("button", { name: "Show Details" }).click();

    await waitFor(() => {
      expect(screen.queryByText(/(\d|no) headers/i)).toBeInTheDocument();
      expect(screen.queryByText(/(\d|no) params/i)).toBeInTheDocument();
    });
  });
});

describe("Filter", () => {
  test("has option to show ALL", () => {
    render(
      <MockedProvider mocks={mocks} addTypename={false}>
        <Requests filters={["PUT", "HEAD"]} />
      </MockedProvider>
    );

    expect(screen.getByRole("button", { name: "ALL" })).toBeInTheDocument();
  });

  test("shows list of standard methods", () => {
    render(
      <MockedProvider mocks={mocks} addTypename={false}>
        <Requests filters={["PUT", "HEAD"]} />
      </MockedProvider>
    );

    expect(screen.getByRole("button", { name: "PUT" })).toBeInTheDocument();
    expect(screen.getByRole("button", { name: "HEAD" })).toBeInTheDocument();
  });

  test("selecting a filter updates button text", async () => {
    render(
      <MockedProvider mocks={mocks} addTypename={false}>
        <Requests filters={["PUT", "HEAD"]} />
      </MockedProvider>
    );

    expect(
      screen.getByRole("button", { name: "Filter: ALL" })
    ).toBeInTheDocument();
    expect(screen.getByRole("button", { name: "PUT" })).toBeInTheDocument();

    screen.getByRole("button", { name: "PUT" }).click();

    await waitFor(() => {
      expect(
        screen.queryByRole("button", { name: "Filter: PUT" })
      ).toBeInTheDocument();
      expect(
        screen.queryByRole("button", { name: "Filter: ALL" })
      ).not.toBeInTheDocument();
    });
  });

  test("selecting a filter filters requests", async () => {
    render(
      <MockedProvider mocks={mocks} addTypename={false}>
        <Requests filters={["PUT", "HEAD", "GET"]} />
      </MockedProvider>
    );

    await waitFor(() => new Promise((resolve) => setTimeout(resolve, 0)));

    expect(
      screen.getByRole("heading", { name: "1 Request" })
    ).toBeInTheDocument();

    screen.getByRole("button", { name: "PUT" }).click();

    await expect(
      screen.getByRole("heading", { name: "0 Requests" })
    ).toBeInTheDocument();

    screen.getByRole("button", { name: "GET" }).click();

    await expect(
      screen.getByRole("heading", { name: "1 Request" })
    ).toBeInTheDocument();
  });
});
