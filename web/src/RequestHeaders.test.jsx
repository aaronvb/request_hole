import { expect, describe, test } from "vitest";
import { render, screen } from "@testing-library/react";
import "@testing-library/jest-dom";
import RequestHeaders from "./RequestHeaders";

const headers = {
  Accept: ["*/*"],
  "Accept-Encoding": ["gzip, deflate"],
  "Accept-Language": ["en-US,en;q=0.5"],
  Connection: ["keep-alive"],
  "Content-Length": ["17"],
  "Content-Type": ["application/json"],
  Origin: ["http://localhost:3000"],
  Referer: ["http://localhost:3000/"],
};

describe("RequestHeaders", () => {
  test("renders count", () => {
    render(<RequestHeaders headers={headers} />);

    expect(screen.getByText(/8 headers/i)).toBeInTheDocument();
  });

  test("renders no count if nil", () => {
    render(<RequestHeaders />);

    expect(screen.getByText(/0 headers/i)).toBeInTheDocument();
  });

  test.each(Object.keys(headers))("renders header %s", (header) => {
    render(<RequestHeaders headers={headers} />);

    expect(screen.getByText(header)).toBeInTheDocument();
    expect(screen.getByText(headers[header])).toBeInTheDocument();
  });
});
