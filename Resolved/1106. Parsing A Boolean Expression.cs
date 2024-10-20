// https://leetcode.com/problems/parsing-a-boolean-expression/
public class Solution {
    public bool ParseBoolExpr(string expression) {
        var i = 0;
        return ParseLexem(expression, ref i);                
    }

    private bool ParseFunction(string expression, ref int index)
    {
        var functionType = expression[index];
        index++; // read function type
        ReadSymbol(expression, '(', ref index);
        var parameters = new List<bool>();
        var first = ParseLexem(expression, ref index); 
        parameters.Add(first);
        while (expression[index] != ')')
        {
            ReadSymbol(expression, ',', ref index);
            var exprResult = ParseLexem(expression, ref index);
            parameters.Add(exprResult);
        }
        index++;
        return EvaluateFunction(functionType, parameters);
    }

    private bool EvaluateFunction(char functionType, List<bool> parameters)
    {
        if (functionType == '!')
        {
            return !parameters[0];
        }
        var result = parameters[0];
        for (var i = 1; i < parameters.Count; i++)
        {
            var currParam = parameters[i];
            if (functionType == '&')
            {
                result = result && currParam;
            }
            else 
            {
                result = result || currParam;
            }
        }
        return result;
    }

    private void ReadSymbol(string expression, char symbol, ref int index)
    {
        if (expression[index] != symbol)
        {
            throw new ArgumentException($"Expected {symbol}, received {expression[index]}");
        }
        index++;
    }

    private bool ParseLexem(string expression, ref int index) 
    {
        var first = expression[index];
        if (first == 't')
        {
            index++;
            return true;
        }
        if (first == 'f')
        {
            index++;
            return false;
        }
        return ParseFunction(expression, ref index);
    }
}
