import { useState } from "react";

import LoginForm from "./components/LoginForm";
import MatrixForm from "./components/MatrixForm";
import MatrixTable from "./components/MatrixTable";
import StatisticsCard from "./components/StatisticsCard";

import { computeQR, login } from "./services/api";
import type { QRResponse } from "./types/matrix";

function App() {
  const [token, setToken] = useState("");
  const [matrixText, setMatrixText] = useState("1,2\n3,4\n5,6");

  const [result, setResult] = useState<QRResponse | null>(null);

  const [error, setError] = useState("");

  async function handleLogin() {
    try {
      const jwt = await login("admin", "admin123");

      setToken(jwt);
      setError("");
    } catch {
      setError("Login failed");
    }
  }

  async function handleCompute() {
    try {
      const matrix = matrixText
        .trim()
        .split("\n")
        .map((row) => row.split(",").map((value) => Number(value.trim())));

      const response = await computeQR(token, {
        matrix,
      });

      setResult(response);
      setError("");
    } catch {
      setError("Request failed");
    }
  }

  return (
    <div>
      <h1>Matrix QR Platform</h1>

      {!token ? (
        <LoginForm onLogin={handleLogin} />
      ) : (
        <>
          <p>Authenticated</p>

          <MatrixForm
            matrixText={matrixText}
            onChange={setMatrixText}
            onSubmit={handleCompute}
          />
        </>
      )}

      {error && <p>{error}</p>}

      {result && (
        <>
          <MatrixTable title="Q Matrix" matrix={result.q} />

          <MatrixTable title="R Matrix" matrix={result.r} />

          <StatisticsCard statistics={result.statistics} />
        </>
      )}
    </div>
  );
}

export default App;
