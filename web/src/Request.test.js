import { render, screen, waitFor } from "@testing-library/react";
import Request from "./Request";

describe("Request", () => {
  test("renders url", () => {
    render(<Request fields={{ url: "/foobar" }} />);

    expect(screen.getByText("/foobar")).toBeInTheDocument();
  });

  test("renders method", () => {
    render(<Request fields={{ method: "POST" }} />);

    expect(screen.getByText("POST")).toBeInTheDocument();
  });

  test("renders created_at time", () => {
    render(<Request fields={{}} created_at={"2000-01-01"} />);

    expect(
      screen.getByText(/(seconds|minutes|hours|days|weeks|months|years) ago/i)
    ).toBeInTheDocument();
  });
});

describe("Details", () => {
  test("clicking hide should hide the details", async () => {
    render(
      <Request
        param_fields={{}}
        headers={{}}
        showAllDetails={true}
        fields={{ url: "/foobar" }}
      />
    );

    expect(screen.getByText(/(\d|no) headers/i)).toBeInTheDocument();
    expect(screen.getByText(/(\d|no) params/i)).toBeInTheDocument();

    screen.getByTestId("toggleDetails").click();

    await waitFor(() => {
      expect(screen.queryByText(/(\d|no) headers/i)).not.toBeInTheDocument();
      expect(screen.queryByText(/(\d|no) params/i)).not.toBeInTheDocument();
    });
  });

  test("clicking show should show the details", async () => {
    render(
      <Request
        param_fields={{}}
        headers={{}}
        showAllDetails={false}
        fields={{ url: "/foobar" }}
      />
    );

    expect(screen.queryByText(/(\d|no) headers/i)).not.toBeInTheDocument();
    expect(screen.queryByText(/(\d|no) params/i)).not.toBeInTheDocument();

    screen.getByTestId("toggleDetails").click();

    await waitFor(() => {
      expect(screen.queryByText(/(\d|no) headers/i)).toBeInTheDocument();
      expect(screen.queryByText(/(\d|no) params/i)).toBeInTheDocument();
    });
  });
});
