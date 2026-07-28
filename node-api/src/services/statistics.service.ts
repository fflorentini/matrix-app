import { QRRequest, StatisticsResponse } from "../types/matrix";

export function calculateStatistics(data: QRRequest): StatisticsResponse {
  const values = [...data.q.flat(), ...data.r.flat()];

  const sum = values.reduce((acc, value) => acc + value, 0);

  const average = values.length > 0 ? sum / values.length : 0;

  const max = Math.max(...values);
  const min = Math.min(...values);

  const hasDiagonalMatrix = isDiagonal(data.q) || isDiagonal(data.r);

  return {
    max,
    min,
    average,
    sum,
    hasDiagonalMatrix,
  };
}

function isDiagonal(matrix: number[][]): boolean {
  const rows = matrix.length;

  if (rows === 0) {
    return false;
  }

  const cols = matrix[0].length;

  if (rows !== cols) {
    return false;
  }

  for (let i = 0; i < rows; i++) {
    for (let j = 0; j < cols; j++) {
      if (i !== j && matrix[i][j] !== 0) {
        return false;
      }
    }
  }

  return true;
}
