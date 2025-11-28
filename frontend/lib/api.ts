// frontend/lib/api.ts
import axios from 'axios';

// 1. Get the base URL from the environment
const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

// 2. Define the response type we expect from Go
interface CalculationResponse {
    result: number;
    error?: string;
}

// 3. The reusable function
export const calculate = async (
    operation: 'add' | 'sub' | 'mul' | 'div' | 'mod',
    num1: number,
    num2: number
): Promise<number> => {
    try {
        // This calls, for example: http://localhost:8080/add
        const response = await axios.post<CalculationResponse>(`${API_URL}/${operation}`, {
            num1: num1,
            num2: num2,
        });

        return response.data.result;
    } catch (error: any) {
        // If the backend returns a specific error message (like "division by zero")
        if (error.response && error.response.data && error.response.data.error) {
            throw new Error(error.response.data.error);
        }
        throw new Error('Network error or server unavailable');
    }
};