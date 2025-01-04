-- AlterTable
ALTER TABLE "accounts" ALTER COLUMN "createdAt" SET DEFAULT CURRENT_TIMESTAMP,
ALTER COLUMN "updatedAt" DROP NOT NULL;

-- AlterTable
ALTER TABLE "projects" ADD COLUMN     "description" TEXT;

-- CreateTable
CREATE TABLE "session_token" (
    "token" TEXT NOT NULL,
    "accountId" UUID NOT NULL,
    "revoke" BOOLEAN NOT NULL DEFAULT false,
    "expirationTime" TIMESTAMPTZ NOT NULL,
    "createdAt" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),

    CONSTRAINT "session_token_pkey" PRIMARY KEY ("token")
);

-- AddForeignKey
ALTER TABLE "session_token" ADD CONSTRAINT "session_token_accountId_fkey" FOREIGN KEY ("accountId") REFERENCES "accounts"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
