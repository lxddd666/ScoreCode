# Introduction

This is a python exe file to convert "locale.xlsx" into json files.

# Remarks

1. Files required:

   - locale.xlsx (utils/excel-to-json)
   - localeList.txt (utils/excel-to-json)
   - excel-to-json.exe
   - {locale}.json (utils/excel-to-json/json)

2. Locale Google Spreadsheet link: https://docs.google.com/spreadsheets/d/1HvYQ-6itzDxEUv7IHOU1hHffHXevnrLLouK1JAgKhoc/edit#gid=0

3. Update localeList.txt if there are more languages added in the future
   - The default locales are en, zh_cn & zh_tw.

# Getting Started

1. Insert the contents needed in Google Spreadsheet.

2. Download the Google Spreadsheet as locale.xlsx.

3. Replace the locale.xlsx under "histudio/utils/excel-to-json".

4. Check if the locales in localeList are correct.

5. Run the excel-to-json.exe file.

6. Replace the OLD json files under "web/src/utils/locales" with the NEW json files under "utils/excel-to-json/json".
