interface MatrixFormProps {
  matrixText: string;
  onChange: (value: string) => void;
  onSubmit: () => void;
}

export default function MatrixForm({
  matrixText,
  onChange,
  onSubmit,
}: MatrixFormProps) {
  return (
    <>
      <textarea
        rows={6}
        cols={40}
        value={matrixText}
        onChange={(e) => onChange(e.target.value)}
      />

      <br />
      <br />

      <button onClick={onSubmit}>Compute QR</button>
    </>
  );
}
