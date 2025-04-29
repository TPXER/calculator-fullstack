'use client'

import { useState } from 'react'

export default function Home() {
  const [a, setA] = useState<number | ''>('')
  const [b, setB] = useState<number | ''>('')
  const [op, setOp] = useState<string>('add')
  const [result, setResult] = useState<string | null>(null)
  const [error, setError] = useState<string | null>(null)

  const handleCalc = async () => {
    setError(null)
    setResult(null)

    if (a === '' || b === '') {
      setError("请输入完整数字")
      return
    }

    try {
      const res = await fetch(`http://localhost:8080/calculator.CalculatorService/${op}`, {
        method: 'POST',
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ a, b }),
      })
      const json = await res.json()
      if (json && typeof json.result === 'number') {
        setResult(json.result.toString())
      } else {
        setError("后端响应异常")
      }
    } catch {
      setError("NetworkError：无法连接服务")
    }
  }

  const handleClear = () => {
    setA('')
    setB('')
    setResult(null)
    setError(null)
  }

  return (
    <div className="min-h-screen bg-neutral-900 text-white flex flex-col justify-center items-center px-4">
      <h1 className="text-6xl font-bold mb-10 font-sans">Calculator</h1>

      <div className="w-full max-w-xl flex items-center justify-between bg-neutral-800 border border-gray-700 rounded-full px-4 py-3 shadow-lg">
        <input
          type="number"
          value={a}
          onChange={(e) => setA(Number(e.target.value))}
          placeholder="数字 A"
          className="w-1/4 bg-transparent outline-none text-lg text-white placeholder-gray-400"
        />

        <select
          value={op}
          onChange={(e) => setOp(e.target.value)}
          className="w-1/4 bg-transparent outline-none text-lg text-white text-center"
        >
          <option value="add">+</option>
          <option value="subtract">−</option>
          <option value="multiply">×</option>
          <option value="divide">÷</option>
        </select>

        <input
          type="number"
          value={b}
          onChange={(e) => setB(Number(e.target.value))}
          placeholder="数字 B"
          className="w-1/4 bg-transparent outline-none text-lg text-white placeholder-gray-400"
        />
      </div>

      <div className="mt-6 flex gap-4">
        <button
          onClick={handleCalc}
          className="bg-blue-600 hover:bg-blue-700 text-white font-medium px-6 py-2 rounded-full transition"
        >
          计算
        </button>
        <button
          onClick={handleClear}
          className="bg-gray-700 hover:bg-gray-600 text-white font-medium px-6 py-2 rounded-full transition"
        >
          清空
        </button>
      </div>

      {result && (
        <div className="mt-8 text-green-400 text-xl font-semibold">
          ✅ 结果：{result}
        </div>
      )}

      {error && (
        <div className="mt-8 text-red-400 text-xl font-semibold">
          ❌ {error}
        </div>
      )}

      {/* 页脚 */}
      <footer className="absolute bottom-0 w-full text-sm text-gray-500 text-center py-4 border-t border-neutral-800">
        <div className="flex justify-center gap-4">
          <a href="#" className="hover:underline">关于</a>
          <a href="#" className="hover:underline">隐私</a>
          <a href="#" className="hover:underline">设置</a>
        </div>
      </footer>
    </div>
  )
}
