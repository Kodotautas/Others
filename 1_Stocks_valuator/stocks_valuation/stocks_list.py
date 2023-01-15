# Description: Get stocks list from Finviz (based on filters) and save it to csv file.
import os
from finviz.screener import Screener
import pandas as pd

cwd = os.getcwd()
print(cwd)

# --------------------------------- GET DATA --------------------------------- #
filters = ['sec_technology'] #tech stocks
stock_list = Screener(filters=filters, table='Valuation', order='price')  # Get the performance table and sort it by price ascending
# stock_list = Screener(table='Valuation', order='price')  # Get the performance table and sort it by price ascending

# Add more filters
stock_list.add(filters=['cap_smallover', 'fa_estltgrowth_pos'])  # Show stocks with hight market cap and positive earnings growth
stock_list.to_csv(f'{cwd}/data/stocks.csv')  # Export the table to a csv file

print('Stocks list saved')