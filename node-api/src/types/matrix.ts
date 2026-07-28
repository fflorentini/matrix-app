export interface QRRequest {
  q: number[][];
  r: number[][];
}

export interface StatisticsResponse {
  max: number;
  min: number;
  average: number;
  sum: number;
  hasDiagonalMatrix: boolean;
}
