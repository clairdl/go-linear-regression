# why?
- i wanted to kill two birds with one stone and do **machine learning** in **golang**
- at least now i understand [why we use python](https://pandas.pydata.org/) for data science :D

## Writeup (incomplete)

### Regression

Ah, regression. The statistical process of modelling relationships between a dependent variables and independent variable(s), enabling us to predict new values. Note that regression techniques are generally concerned with predicting continous values, as opposed to a discreet set of categories.

### Linear regression

This brings us to possibly the most fundamental model, linear regression, expressed with the battle tested equation:

`y = mx + b`

Which describes a line with gradient `m` and y-intercept `b`.

One way of actually computing `m` and `b` is with the [ordinary least squares](https://en.wikipedia.org/wiki/Ordinary_least_squares) method:

1. Randomise values for both m and b to create an example line
2. Find the distance between the example line and each value in the dataset (these distances are called 'errors'):

<img src="./docs/ols.png">

3. Sum the squares of these errors:

<p align="center">
<img src="./docs/eq0.png">
</p>

Now, we iteratively adjust the values of `m` and `b` in order to minimize this sum. A ubiquitous optimization technique to find the local minima is called _gradient descent_, but that's a topic for another day :)

The accuracy and performance of linear regression is dependant on its **assumptions**:

- **Linearity**: there is a linear relationship between the dependant variable and the independant variable(s)
- **Normality**: your variables are distributed normally
- **No multicollinearity**: your independant variables should not be predictors of eachother, almost by definition
- **No auto-correlation**: a fancy way of saying your variables should not depend on themselves, i.e they are not values in a time series; Tesla's share price, for example

**Pitfalls**:
- Extrapolation beyond the model can quickly become very inaccurate
- Extreme outliers can throw off the model
