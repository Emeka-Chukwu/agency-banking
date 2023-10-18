-- -- Enable the UUID extension if it's not already enabled
-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- -- Create Agents Table with UUID as the primary key
-- CREATE TABLE Agents (
--     ID UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
--     Name VARCHAR(255) NOT NULL,
--     Email VARCHAR(255) NOT NULL UNIQUE,
--     Password VARCHAR(255) NOT NULL,
--     Contact_Info VARCHAR(255) NOT NULL
-- );

-- -- Create Agent_Accounts Table with UUID as the primary key
-- CREATE TABLE Accounts (
--     ID UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
--     Agent_ID UUID NOT NULL,
--     Balance DECIMAL(10, 2) DEFAULT 0.0,
--     FOREIGN KEY (Agent_ID) REFERENCES Agents(Agent_ID)
-- );

-- -- Create POS_Terminals Table with UUID as the primary key
-- CREATE TABLE POS_Terminals (
--     ID UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
--     Serial_Number VARCHAR(255) NOT NULL,
--     Agent_ID UUID NOT NULL,
--     Status ENUM('active', 'active',"deactivated") NOT NULL,
--     Other_Properties VARCHAR(255),
--     FOREIGN KEY (Agent_ID) REFERENCES Agents(Agent_ID)
-- );

-- -- Create Transactions Table with UUID as the primary key
-- CREATE TABLE Transactions (
--     ID UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
--     Type VARCHAR(255) NOT NULL,
--     Amount DECIMAL(10, 2) NOT NULL,
--     Timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     Status ENUM('success', 'pending', 'failed') NOT NULL,
--     Agent_ID UUID NOT NULL,
--     Transaction_Description VARCHAR(255),
--     FOREIGN KEY (Agent_ID) REFERENCES Agents(Agent_ID)
-- );

-- -- Create Charges Table with UUID as the primary key
-- CREATE TABLE Charges (
--     ID UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
--     Transaction_Type VARCHAR(255) NOT NULL,
--     Percentage DECIMAL(5, 2) NOT NULL
-- );

-- -- Create Company_Account Table with UUID as the primary key
-- CREATE TABLE Company_Account (
--     Account_ID UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
--     Balance DECIMAL(10, 2) DEFAULT 0.0 NOT NULL
-- );

-- -- Create Transaction_Processing Table with UUID as the primary key
-- CREATE TABLE Transaction_Processing (
--     ID UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
--     Transaction_ID UUID NOT NULL,
--     Transaction_Reference VARCHAR(255) NOT NULL,
--     Status ENUM('pending', 'completed','failed') NOT NULL,
--     FOREIGN KEY (Transaction_ID) REFERENCES Transactions(Transaction_ID)
-- );








-- delete from schema_migrations;

-- Enable the UUID extension if it's not already enabled
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create Agents Table with UUID as the primary key
CREATE TABLE Agents (
    ID UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    Name VARCHAR(255) NOT NULL,
    Email VARCHAR(255) NOT NULL UNIQUE,
    Password VARCHAR(255) NOT NULL,
    Contact_Info VARCHAR(255) NOT NULL
);

-- Create Agent_Accounts Table with UUID as the primary key
CREATE TABLE Accounts (
    ID UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    Agent_ID UUID NOT NULL,
    Balance DECIMAL(10, 2) DEFAULT 0.0,
    FOREIGN KEY (Agent_ID) REFERENCES Agents(ID)
);

-- Create POS_Terminals Table with UUID as the primary key
CREATE TABLE POS_Terminals (
    ID UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    Serial_Number VARCHAR(255) NOT NULL,
    Agent_ID UUID NOT NULL,
    Status VARCHAR(255) NOT NULL,
    FOREIGN KEY (Agent_ID) REFERENCES Agents(ID)
);

-- Create Transactions Table with UUID as the primary key
CREATE TABLE Transactions (
    ID UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    Type VARCHAR(255) NOT NULL,
    Amount DECIMAL(10, 2) NOT NULL,
    Timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    Status VARCHAR(255) NOT NULL,
    Agent_ID UUID NOT NULL,
    Transaction_Description VARCHAR(255),
    FOREIGN KEY (Agent_ID) REFERENCES Agents(ID)
);

-- Create Charges Table with UUID as the primary key
CREATE TABLE Charges (
    ID UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    Transaction_Type VARCHAR(255) NOT NULL,
    Percentage DECIMAL(5, 2) NOT NULL
);

-- Create Company_Account Table with UUID as the primary key
CREATE TABLE Company_Account (
    Account_ID UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    Balance DECIMAL(10, 2) DEFAULT 0.0 NOT NULL
);

-- Create Transaction_Processing Table with UUID as the primary key
CREATE TABLE Transaction_Processing (
    ID UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    Transaction_ID UUID NOT NULL,
    Transaction_Reference VARCHAR(255) NOT NULL,
    Status varchar(255) NOT NULL,
    FOREIGN KEY (Transaction_ID) REFERENCES Transactions(ID)
);

