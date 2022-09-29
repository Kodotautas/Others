import os
from finviz.screener import Screener
import pandas as pd

cwd = os.getcwd()

# --------------------------------- GET DATA --------------------------------- #
filters = ['sec_technology'] #sec_consumercyclical
stock_list = Screener(filters=filters, table='Valuation', order='price')  # Get the performance table and sort it by price ascending

# Add more filters
stock_list.add(filters=['cap_midover', 'fa_estltgrowth_pos'])  # Show stocks with high dividend yield
stock_list.to_csv(f'{cwd}/data/stocks.csv')

print('Stocks list saved')