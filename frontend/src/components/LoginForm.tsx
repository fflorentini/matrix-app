interface LoginFormProps {
  onLogin: () => void;
}

export default function LoginForm({ onLogin }: LoginFormProps) {
  return <button onClick={onLogin}>Login</button>;
}
