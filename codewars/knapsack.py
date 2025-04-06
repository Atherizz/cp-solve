def knapsack(items, w_limit):
    weights = []
    values = []
    
    for item in items:
        weights.append(item[0])
        values.append(item[1])
    

    n = len(weights)
    

    table = []
    for i in range(n + 1):
        row = []
        for j in range(w_limit + 1):
            row.append(0)
        table.append(row)
    
    maks = 0
    

    for i in range(1, n + 1):
        for j in range(1, w_limit + 1):
            if i == 1 and j >= weights[i-1]:
                table[i][j] = values[i-1]
            else:
                if j < weights[i-1]:
                    table[i][j] = table[i-1][j]
                else:
                    a = values[i-1] + table[i-1][j - weights[i-1]]
                    b = table[i-1][j]
                    if a > b:
                        maks = a
                    else:
                        maks = b
                    table[i][j] = maks
    
    max_profit = 0
    chosen_values = []
    i = n
    j = w_limit
    
    while i > 0 and j > 0:
        if table[i][j] != table[i-1][j]:
            chosen_values.append(values[i-1])
            max_profit += values[i-1]
            j -= weights[i-1]
        i -= 1
    
    if i == 0 and j >= weights[0] and table[0][j] > 0:
        chosen_values.append(values[0])
        max_profit += values[0]
    
    
    return [max_profit, [sorted(chosen_values)]]
