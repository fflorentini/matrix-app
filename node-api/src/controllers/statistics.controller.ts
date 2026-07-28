import { Request, Response } from "express";

import { QRRequest } from "../types/matrix";
import { calculateStatistics } from "../services/statistics.service";

export function getStatistics(req: Request, res: Response): void {
  const body = req.body as QRRequest;

  const result = calculateStatistics(body);

  res.json(result);
}
