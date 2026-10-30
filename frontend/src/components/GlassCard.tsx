export default function GlassCard({ children, className = '' }: { children: React.ReactNode; className?: string }) {
  return (
    <div className={`glass rounded-3xl p-5 ${className}`}>
      {children}
    </div>
  )
}
