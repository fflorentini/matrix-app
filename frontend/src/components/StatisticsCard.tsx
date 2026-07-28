import type { Statistics } from "../types/matrix";

interface StatisticsCardProps {
  statistics: Statistics;
}

export default function StatisticsCard({ statistics }: StatisticsCardProps) {
  return (
    <>
      <h2>Statistics</h2>

      <pre>{JSON.stringify(statistics, null, 2)}</pre>
    </>
  );
}
