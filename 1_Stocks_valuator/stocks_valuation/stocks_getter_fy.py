import yfinance as yf
import pandas as pd
import os
from datetime import date
from stocks_list import filters

cwd = os.getcwd()
sector = filters[0] #get loaded sector
today = date.today()


# -------------------------------- STOCKS LIST ------------------------------- #
stocks_df = pd.read_csv(f'{cwd}/data/stocks.csv')
stocks = stocks_df['Ticker'].to_list()
stocks = ['MSFT', 'NVDA']
print('Stocks list loaded!')


# ------------------------------- GET INFO TAB ------------------------------- #
print('Getting stocks info tab...')
tickerStrings = stocks
count = 0
df_list = list()
for ticker in tickerStrings:
    try:
        data =  yf.Ticker(ticker).info
        data['ticker'] = ticker  # add this column because the dataframe doesn't contain a column with the ticker
        df_list.append(data)
        count += 1
        print(ticker, 'info loaded ', str(int(count / len(stocks)*100))+'%')
    except TypeError:
        print("TypeError accured. Could be wrong ticker.")

# combine all dataframes into a single dataframe
df_info = pd.DataFrame(df_list)
df_info = df_info[['ticker', 'industry', 'marketCap', 'regularMarketPrice', '52WeekChange', 'dividendYield']]

df_info['dividendYield'] = df_info['dividendYield'] * 100 #convert to percentage
df_info['marketCap'] = df_info['marketCap'] / 1000000000 #convert to billions
print('Stocks info loaded.')


# ----------------------------- GET ANALYSIS TAB ----------------------------- #
print('Getting stocks analysis tab...')
tickerStrings = stocks
count = 0
df_list_b = list()
for ticker in tickerStrings:
    try:
        data_b =  yf.Ticker(ticker).analysis
        data_b['ticker'] = ticker  # add this column because the dataframe doesn't contain a column with the ticker
        df_list_b.append(data_b)
        count += 1
        print(ticker, 'analysis loaded ', str(int(count / len(stocks)*100))+'%')
    except TypeError:
        print("TypeError accured. Could be wrong ticker.")

# combine all dataframes into a single dataframe
df_analysis = pd.concat(df_list_b).reset_index()

#filter columns and leave only positive earnings
df_analysis = df_analysis[['Period', 'Growth', 'Earnings Estimate Avg', 'ticker']]
df_analysis['Growth'] = df_analysis['Growth'] * 100

#shift one row down
df_analysis['Earnings Estimate Avg'] = df_analysis['Earnings Estimate Avg'].shift(1)

#filter only needed periods
df_analysis = df_analysis[df_analysis['Period'] == '+5Y']
df_analysis = df_analysis.drop(['Period'], axis=1)

#merge dataframes and drop nan
df = df_info.merge(df_analysis, how='left', on='ticker')
print('Stocks analysis loaded.')


# ----------------------------- STOCKS VALUATION ----------------------------- #
#calcualte fair value and MOS
df['fair_value'] = (df['Growth']/2 + 8) * df['Earnings Estimate Avg']
df['MOS'] = (df['fair_value'] - df['regularMarketPrice']) / df['regularMarketPrice']

#filter negative earnings and MOS
df = df[(df['Earnings Estimate Avg'] > 0) & (df['MOS']>0)]

#sort by Margin of Safety
df = df.sort_values(by=['MOS'], ascending=False)

#export csv
df.to_excel(f'{cwd}/data/stocks_valuation_{sector}_{today}.xlsx')
print('Stocks values calculated successfully and exported!')