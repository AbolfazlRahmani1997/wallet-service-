-- Create TABLE Wallets
-- (
--     Id        varchar(255) primary key,
--     Balance   decimal(8, 2),
--     CreatedAt timestamp,
--     UpdatedAt timestamp
--
-- );
-- CREATE TYPE documentType as ENUM (
--     'CASH_IN','CASH_OUT','TRANSFER'
--     );
-- CREATE TABLE Documents
-- (
--     Id                varchar(255) primary key,
--     WalletOrigin      varchar(255),
--     WalletDestination varchar(255),
--     GasFee            decimal(8, 2),
--     TrackingCode      varchar(255),
--     DocumentType      documentType
--
-- );
--
-- CREATE TYPE transactionType as ENUM (
--     'DEPOSIT','WITHDRAW','FEE'
--     );
-- CREATE TABLE Transactions
-- (
--     Id              varchar(255) primary key,
--     WalletId        varchar(255),
--     DocumentId      varchar(255),
--     Amount          decimal(8, 2),
--     TransactionType transactionType
-- );


CREATE FUNCTION update_wallet_balance()
    RETURNS TRIGGER
AS $$
BEGIN
    UPDATE Wallets
    SET Balance = (SELECT sum(Amount) FROM transactions WHERE wallet_id = NEW.wallet_id)
    WHERE Id = NEW.wallet_id;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
Create TRIGGER wallet_balance
    AFTER Insert or UPDATE
    ON transactions
    FOR EACH ROW
EXECUTE PROCEDURE update_wallet_balance();


