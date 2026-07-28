import express, { Request, Response } from "express";

import { getStatistics } from "./controllers/statistics.controller";

const app = express();

app.use(express.json());

app.get("/health", (_req: Request, res: Response) => {
  res.json({
    status: "ok",
  });
});

app.post("/api/statistics", getStatistics);

const PORT = 3000;

app.listen(PORT, () => {
  console.log(`Node API listening on ${PORT}`);
});
