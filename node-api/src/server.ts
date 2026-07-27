import express, { Request, Response } from "express";

const app = express();

app.get("/health", (_req: Request, res: Response) => {
  res.json({
    status: "ok",
  });
});

const PORT = 3000;

app.listen(PORT, () => {
  console.log(`Node API listening on ${PORT}`);
});
