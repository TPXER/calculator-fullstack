'use client';
import { useState } from 'react';

export default function Home() {
  // 定义组件状态
  const [numA, setNumA] = useState('');
  const [numB, setNumB] = useState('');
  const [op, setOp] = useState('Add');    // 运算符，默认加法
  const [result, setResult] = useState(null);
  const [error, setError] = useState('');

  // 处理表单提交
  const handleCalculate = async () => {
    setError(''); // 重置错误
    setResult(null);
    if (numA === '' || numB === '') {
      setError('请输入两个数字');
      return;
    }
    // 准备请求数据
    const a = parseFloat(numA);
    const b = parseFloat(numB);
    const url = `http://localhost:8080/calculator.CalculatorService/${op}`;
    try {
      const res = await fetch(url, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ a, b })
      });
      if (!res.ok) {
        // HTTP 错误，从响应中提取错误信息
        const errData = await res.json();
        const errMsg = errData.error?.message || '请求失败';
        throw new Error(errMsg);
      }
      const data = await res.json();
      setResult(data.result);
    } catch (err) {
      setError(`计算出错：${err.message}`);
    }
  };

  return (
    <div style={{ maxWidth: '400px', margin: '50px auto', fontFamily: 'sans-serif' }}>
      <h1>在线计算器</h1>
      <div style={{ marginBottom: '8px' }}>
        <input
          type="number"
          placeholder="数字 A"
          value={numA}
          onChange={(e) => setNumA(e.target.value)}
          style={{ width: '100px' }}
        />
        <select value={op} onChange={(e) => setOp(e.target.value)} style={{ margin: '0 8px' }}>
          <option value="Add">＋ 加法</option>
          <option value="Subtract">− 减法</option>
          <option value="Multiply">× 乘法</option>
          <option value="Divide">÷ 除法</option>
        </select>
        <input
          type="number"
          placeholder="数字 B"
          value={numB}
          onChange={(e) => setNumB(e.target.value)}
          style={{ width: '100px' }}
        />
      </div>
      <button onClick={handleCalculate}>计算</button>
      <div style={{ marginTop: '16px', minHeight: '1.5em' }}>
        {error ? <span style={{ color: 'red' }}>{error}</span> : null}
        {result !== null ? <span>结果：<b>{result}</b></span> : null}
      </div>
    </div>
  );
}
