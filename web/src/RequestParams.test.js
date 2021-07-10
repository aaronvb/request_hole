import { render, screen, waitFor } from "@testing-library/react";
import RequestParams from "./RequestParams";

const params = {
  foo: "bar",
  hello: "world",
  aloha: "friday",
};

describe("RequestParams", () => {
  describe("k/v params", () => {
    test("renders query count", () => {
      render(<RequestParams params={{ query: params }} />);

      expect(screen.getByText(/3 query params/i)).toBeInTheDocument();
    });

    test("renders form count", () => {
      render(<RequestParams params={{ form: params }} />);

      expect(screen.getByText(/3 form params/i)).toBeInTheDocument();
    });

    test.each(Object.keys(params))("renders query param %s", (param) => {
      render(<RequestParams params={{ query: params }} />);

      expect(screen.getByText(param)).toBeInTheDocument();
      expect(screen.getByText(params[param])).toBeInTheDocument();
    });

    test.each(Object.keys(params))("renders form param %s", (param) => {
      render(<RequestParams params={{ form: params }} />);

      expect(screen.getByText(param)).toBeInTheDocument();
      expect(screen.getByText(params[param])).toBeInTheDocument();
    });
  });

  describe("json", () => {
    test("renders json", () => {
      render(<RequestParams params={{ json: params }} />);

      expect(screen.getByText(/json body/i)).toBeInTheDocument();
    });

    test("renders json param", () => {
      render(<RequestParams params={{ json: params }} />);

      expect(screen.getByText(/foo/)).toBeInTheDocument();
      expect(screen.getByText(/bar/)).toBeInTheDocument();
      expect(screen.getByText(/world/)).toBeInTheDocument();
      expect(screen.getByText(/friday/)).toBeInTheDocument();
    });
  });
});
