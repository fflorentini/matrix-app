import type { LoginResponse, QRRequest, QRResponse } from "../types/matrix";

const API_URL = import.meta.env.VITE_API_URL || "http://localhost:8080";

export async function login(
  username: string,
  password: string,
): Promise<string> {
  const response = await fetch(`${API_URL}/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      username,
      password,
    }),
  });

  if (!response.ok) {
    throw new Error("Login failed");
  }

  const data: LoginResponse = await response.json();

  return data.token;
}

export async function computeQR(
  token: string,
  request: QRRequest,
): Promise<QRResponse> {
  const response = await fetch(`${API_URL}/api/qr`, {
    method: "POST",
    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
    body: JSON.stringify(request),
  });

  if (!response.ok) {
    throw new Error("QR request failed");
  }

  return response.json();
}
