# Description: Get stocks list from Finviz (based on filters) and save it to csv file.
import os
from finviz.screener import Screener
import pandas as pd

cwd = os.getcwd()
print('Working directory: ', cwd)

# --------------------------------- GET DATA --------------------------------- #
filters = ['sec_industrials'] #set filter, ['sec_technology'] for tech stocks
stock_list = Screener(filters=filters, table='Valuation', order='price')  # Get the performance table and sort it by price ascending

# Add more filters
stock_list.add(filters=['cap_smallover', 'fa_estltgrowth_pos'])  # Show stocks with hight market cap and positive earnings growth
stock_list.to_csv(f'{cwd}/1_Stocks_valuator/stocks_valuation/data/stocks.csv')  # Export the table to a csv file

print('Stocks list saved')