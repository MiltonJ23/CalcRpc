'use client'; // <--- Essential! Marks this as a Client Component

import { useState } from 'react';
import { calculate } from '@/lib/api'; // Verify this import path matches your setup

export default function CalculatorPage() {
    // --- State Management ---
    const [num1, setNum1] = useState<string>(''); // Keep as string to handle empty inputs
    const [num2, setNum2] = useState<string>('');
    const [result, setResult] = useState<number | null>(null);
    const [error, setError] = useState<string | null>(null);
    const [loading, setLoading] = useState<boolean>(false);

    // --- Logic ---
    const handleOperation = async (op: 'add' | 'sub' | 'mul' | 'div' | 'mod') => {
        // Reset previous states
        setError(null);
        setResult(null);
        setLoading(true);

        // Validate inputs
        if (num1 === '' || num2 === '') {
            setError("Please enter both numbers");
            setLoading(false);
            return;
        }

        try {
            // Call the API layer
            const res = await calculate(op, Number(num1), Number(num2));
            setResult(res);
        } catch (err: any) {
            setError(err.message);
        } finally {
            setLoading(false);
        }
    };

    // --- UI Render ---
    return (
        <main className="min-h-screen flex items-center justify-center bg-gray-50 p-4">
            <div className="bg-white p-8 rounded-xl shadow-lg w-full max-w-md">
                <h1 className="text-3xl font-bold text-center mb-6 text-gray-800">
                    Go Calculator
                </h1>

                {/* Inputs */}
                <div className="space-y-4 mb-6">
                    <div>
                        <label className="block text-sm font-medium text-gray-700 mb-1">Number 1</label>
                        <input
                            type="number"
                            value={num1}
                            onChange={(e) => setNum1(e.target.value)}
                            className="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none text-black"
                            placeholder="0"
                        />
                    </div>
                    <div>
                        <label className="block text-sm font-medium text-gray-700 mb-1">Number 2</label>
                        <input
                            type="number"
                            value={num2}
                            onChange={(e) => setNum2(e.target.value)}
                            className="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none text-black"
                            placeholder="0"
                        />
                    </div>
                </div>

                {/* Operation Buttons */}
                <div className="grid grid-cols-3 gap-3 mb-6">
                    <OperationButton onClick={() => handleOperation('add')} label="+" color="bg-blue-600" />
                    <OperationButton onClick={() => handleOperation('sub')} label="-" color="bg-blue-600" />
                    <OperationButton onClick={() => handleOperation('mul')} label="×" color="bg-blue-600" />
                    <OperationButton onClick={() => handleOperation('div')} label="÷" color="bg-orange-500" />
                    <OperationButton onClick={() => handleOperation('mod')} label="%" color="bg-purple-600" />

                    <button
                        onClick={() => { setNum1(''); setNum2(''); setResult(null); setError(null); }}
                        className="bg-gray-400 hover:bg-gray-500 text-white font-bold py-3 rounded-lg transition"
                    >
                        CLR
                    </button>
                </div>

                {/* Results Area */}
                <div className="bg-gray-100 p-4 rounded-lg min-h-[80px] flex items-center justify-center">
                    {loading ? (
                        <span className="text-gray-500 animate-pulse">Calculating...</span>
                    ) : error ? (
                        <span className="text-red-500 font-medium">{error}</span>
                    ) : result !== null ? (
                        <span className="text-4xl font-bold text-gray-800">{result}</span>
                    ) : (
                        <span className="text-gray-400 text-sm">Result will appear here</span>
                    )}
                </div>
            </div>
        </main>
    );
}

// Helper Component for Buttons to keep code clean
function OperationButton({ onClick, label, color }: { onClick: () => void, label: string, color: string }) {
    return (
        <button
            onClick={onClick}
            className={`${color} hover:opacity-90 text-white text-xl font-bold py-3 rounded-lg transition shadow-md active:scale-95`}
        >
            {label}
        </button>
    );
}