export interface LoginResponse {
  token: string;
}

export interface QRRequest {
  matrix: number[][];
}

export interface Statistics {
  max: number;
  min: number;
  average: number;
  sum: number;
  hasDiagonalMatrix: boolean;
}

export interface QRResponse {
  q: number[][];
  r: number[][];
  statistics: Statistics;
}
