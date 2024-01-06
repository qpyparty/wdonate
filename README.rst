==============================
WDonate
==============================

WDonate is a Go library that provides functionality for working with the wdonate.ru API.

Installation
------------

To install the library, use the following command:

.. code-block:: bash

  go get github.com/qpyparty/wdonate

Usage
-----

Here is an example of using the library to get the balance:

.. code-block:: go

  package main

  import (
      "fmt"
      "github.com/qpyparty/wdonate"
  )

  func main() {
      // Replace <API TOKEN> with your actual API token obtained from your wdonate.ru account
      token := "<API TOKEN>"
      // Replace 0 with your actual VK Bot ID, ensuring it doesn't contain a sign
      botId := 0

      client := wdonate.NewClient(token, botId)

      // Execute a request to get the balance
      response, err := client.GetBalance(wdonate.GetBalanceRequest{})
      if err != nil {
          panic(err)
      }

      fmt.Println(response.Balance)
  }

License
-------

This project is licensed under the MIT License. See the LICENSE file for details.