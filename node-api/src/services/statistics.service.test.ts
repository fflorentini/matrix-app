import { calculateStatistics } from "./statistics.service";

describe("calculateStatistics", () => {
  it("calculates statistics correctly", () => {
    const result = calculateStatistics({
      q: [
        [1, 2],
        [3, 4],
      ],
      r: [
        [5, 6],
        [7, 8],
      ],
    });

    expect(result.max).toBe(8);
    expect(result.min).toBe(1);
    expect(result.sum).toBe(36);
    expect(result.average).toBe(4.5);
    expect(result.hasDiagonalMatrix).toBe(false);
  });

  it("detects a diagonal matrix", () => {
    const result = calculateStatistics({
      q: [
        [1, 0],
        [0, 2],
      ],
      r: [
        [3, 0],
        [0, 4],
      ],
    });

    expect(result.hasDiagonalMatrix).toBe(true);
  });

  it("handles negative values", () => {
    const result = calculateStatistics({
      q: [
        [-1, -2],
        [-3, -4],
      ],
      r: [
        [-5, -6],
        [-7, -8],
      ],
    });

    expect(result.max).toBe(-1);
    expect(result.min).toBe(-8);
    expect(result.sum).toBe(-36);
    expect(result.average).toBe(-4.5);
  });

  it("works when only one matrix is diagonal", () => {
    const result = calculateStatistics({
      q: [
        [1, 2],
        [3, 4],
      ],
      r: [
        [5, 0],
        [0, 6],
      ],
    });

    expect(result.hasDiagonalMatrix).toBe(true);
  });
});
