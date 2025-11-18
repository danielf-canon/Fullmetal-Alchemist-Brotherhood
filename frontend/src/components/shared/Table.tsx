// src/components/shared/Table.tsx
import React from "react";

type Column<T> = { key: string; header: string; render?: (row: T) => React.ReactNode };

export default function Table<T>({ columns, data }: { columns: Column<T>[]; data: T[] }) {
  return (
    <table style={{width: "100%", borderCollapse: "collapse"}}>
      <thead>
        <tr>
          {columns.map(c => <th key={c.key} style={{textAlign: "left", borderBottom: "1px solid #eee", padding: 8}}>{c.header}</th>)}
        </tr>
      </thead>
      <tbody>
        {data.map((row, i) => (
          <tr key={i}>
            {columns.map(c => <td key={c.key} style={{padding: 8, borderBottom: "1px solid #fafafa"}}>{c.render ? c.render(row) : (row as any)[c.key]}</td>)}
          </tr>
        ))}
      </tbody>
    </table>
  );
}
