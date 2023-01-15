import yfinance as yf
import pandas as pd
import os
from datetime import date
from stocks_list import filters

cwd = os.getcwd()
sector = filters[0] #get loaded sector
# sector = 'sec_technology'
today = date.today()


# ----------------------------- LOAD STOCKS LIST ----------------------------- #
stocks_df = pd.read_csv(f'{cwd}/data/stocks.csv')
stocks = stocks_df['Ticker'].to_list()
# stocks = ['MSFT', 'NVDA']
print('Stocks list loaded!')


# ---------------------- GET INFO TAB FROM YAHOO FINANCE --------------------- #
print('Getting stocks info tab from Yahoo Finance...')
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
df_info = df_info[['ticker', 'industry', 'marketCap', 'regularMarketPrice', '52WeekChange', 'dividendYield', 'forwardEps']]

df_info['dividendYield'] = df_info['dividendYield'] * 100 #convert to percentage
df_info['marketCap'] = df_info['marketCap'] / 1000000000 #convert to billions
print('Stocks info loaded.')


# ------------------------------- TRANSFORM DATA ------------------------------ #
leave_columns = ['Ticker', 'Market Cap', 'EPS next Y', 'EPS next 5Y', 'Price']
stocks_df = stocks_df[leave_columns]

# from Market Cap remove B or M character and convert to float
stocks_df['Market Cap'] = stocks_df['Market Cap'].str.replace('B', '')
stocks_df['Market Cap'] = stocks_df['Market Cap'].str.replace('M', '')

# from EPS next Y and EPS next 5Y remove % character and convert to float
stocks_df['EPS next Y'] = stocks_df['EPS next Y'].str.replace('%', '')
stocks_df['EPS next 5Y'] = stocks_df['EPS next 5Y'].str.replace('%', '')

# drop rows where EPS next Y is '-'
stocks_df = stocks_df[stocks_df['EPS next Y'] != '-']

# convert to float
for column in stocks_df.columns[1:]:
    stocks_df[column] = stocks_df[column].astype(float)


# ----------------------------- MERGE DATAFRAMES ----------------------------- #
stocks_df = pd.merge(stocks_df, df_info, how='left', left_on='Ticker', right_on='ticker')


# ----------------------------- STOCKS VALUATION ----------------------------- #
#calcualte fair value and MOS
stocks_df['fair_value'] = (stocks_df['EPS next 5Y']/2 + 5) * stocks_df['forwardEps']
stocks_df['MOS'] = (stocks_df['fair_value'] - stocks_df['Price']) / stocks_df['Price']

#filter negative earnings and MOS
stocks_df = stocks_df[(stocks_df['forwardEps'] > 0) & (stocks_df['MOS']>0)]

#sort by Margin of Safety
stocks_df = stocks_df.sort_values(by=['MOS'], ascending=False)

# reorder columns
stocks_df = stocks_df[['Ticker', 'industry', 'marketCap', 'regularMarketPrice', 'dividendYield', 'Price', 'fair_value', 'MOS']]

#export csv
stocks_df.to_excel(f'{cwd}/data/stocks_valuation_{sector}_{today}.xlsx')
print('Stocks values calculated successfully and exported!')