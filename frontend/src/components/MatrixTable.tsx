interface MatrixTableProps {
  title: string;
  matrix: number[][];
}

export default function MatrixTable({ title, matrix }: MatrixTableProps) {
  return (
    <>
      <h2>{title}</h2>

      <pre>{JSON.stringify(matrix, null, 2)}</pre>
    </>
  );
}
